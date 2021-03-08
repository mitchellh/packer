require('dotenv').config()
const fs = require('fs')
const path = require('path')
const {
  indexContent,
  getDocsSearchObject,
} = require('@hashicorp/react-search/tools')
const resolveNavData = require('../components/remote-plugin-docs/utils/resolve-nav-data')
const fetchGithubFile = require('../components/remote-plugin-docs/utils/fetch-github-file')

// Run indexing
indexContent({ getSearchObjects })

async function getSearchObjects() {
  // Resolve /docs, /guides, and /intro nav data, which
  // corresponds to all the content we will actually render
  // (this avoids indexing non-rendered content, and partials)
  // `docs` content
  const docsNav = await resolveNavData(
    'data/docs-nav-data.json',
    'content/docs',
    { remotePluginsFile: 'data/docs-remote-plugins.json' }
  )
  const docsObjects = await searchObjectsFromNavData(docsNav, 'docs')
  // `guides` content
  const guidesNav = await resolveNavData(
    'data/guides-nav-data.json',
    'content/guides'
  )
  const guidesObjects = await searchObjectsFromNavData(guidesNav, 'guides')
  // `intro` content
  const introNav = await resolveNavData(
    'data/intro-nav-data.json',
    'content/intro'
  )
  const introObjects = await searchObjectsFromNavData(introNav, 'intro')
  // Concat all search objects, and return them
  const searchObjects = [].concat(docsObjects, guidesObjects, introObjects)
  return searchObjects
}

// Given navData, return a flat array of search objects
// for each content file reference in the navData nodes
async function searchObjectsFromNavData(navData, baseRoute = '') {
  const searchObjectsFromNodes = await Promise.all(
    navData.map((n) => searchObjectsFromNavNode(n, baseRoute))
  )
  const flattenedSearchObjects = searchObjectsFromNodes.reduce(
    (acc, searchObjects) => acc.concat(searchObjects),
    []
  )
  return flattenedSearchObjects
}

// Given a navData node, return a flat array of search objects
// for each content file referenced in the node.
// For "leaf" nodes, this will yield an array with a single object.
// For "branch" nodes, this may yield an array with zero or more search objects.
// For all other nodes, this will yield an empty array.
async function searchObjectsFromNavNode(node, baseRoute) {
  // If this is a node, build a search object
  if (node.path) {
    //  Fetch the MDX file content
    const [err, fileString] = node.filePath
      ? //  Read local content from the filesystem
        [null, fs.readFileSync(path.join(process.cwd(), node.filePath), 'utf8')]
      : // Fetch remote content using GitHub's API
        await fetchGithubFile(node.remoteFile)
    if (err) throw new Error(err)
    const searchObject = await getDocsSearchObject(
      path.join(baseRoute, node.path),
      fileString
    )
    return searchObject
  }
  //  If this is a branch, recurse
  if (node.routes) return await searchObjectsFromNavData(node.routes, baseRoute)
  // Otherwise, return an empty array
  // (for direct link nodes, divider nodes)
  return []
}
