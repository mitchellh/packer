set :base_url, "https://www.packer.io/"

activate :breadcrumbs

activate :hashicorp do |h|
  h.name        = "packer"
  h.version     = "0.8.6"
  h.github_slug = "mitchellh/packer"
end
