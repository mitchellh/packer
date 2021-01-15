# This file was autogenerated by the BETA 'packer hcl2_upgrade' command. We
# recommend double checking that everything is correct before going forward. We
# also recommend treating this file as disposable. The HCL2 blocks in this
# file can be moved to other files. For example, the variable blocks could be
# moved to their own 'variables.pkr.hcl' file, etc. Those files need to be
# suffixed with '.pkr.hcl' to be visible to Packer. To use multiple files at
# once they also need to be in the same folder. 'packer inspect folder/'
# will describe to you what is in that folder.

# Avoid mixing go templating calls ( for example ```{{ upper(`string`) }}``` )
# and HCL2 calls (for example '${ var.string_value_example }' ). They won't be
# executed together and the outcome will be unknown.

# See https://www.packer.io/docs/from-1.5/blocks/packer for more info
packer {
  required_version = ">= 1.6.0"
}

# All generated input variables will be of 'string' type as this is how Packer JSON
# views them; you can change their type later on. Read the variables type
# constraints documentation
# https://www.packer.io/docs/from-1.5/variables#type-constraints for more info.
variable "aws_access_key" {
  type      = string
  default   = ""
  sensitive = true
}

variable "aws_region" {
  type = string
}

variable "aws_secondary_region" {
  type    = string
  default = "${env("AWS_DEFAULT_REGION")}"
}

variable "aws_secret_key" {
  type      = string
  default   = ""
  sensitive = true
}

variable "secret_account" {
  type      = string
  default   = "🤷"
  sensitive = true
}

# "timestamp" template function replacement
locals { timestamp = regex_replace(timestamp(), "[- TZ:]", "") }

# amazon-ami data block is generated from your amazon builder source_ami_filter; a data can be referenced in
# source and locals blocks. Read the documentation for source blocks here:
# https://www.packer.io/docs/from-1.5/blocks/data
data "amazon-ami" "autogenerated_1" {
  filters = {
    name                = "ubuntu/images/*/ubuntu-xenial-16.04-amd64-server-*"
    root-device-type    = "ebs"
    virtualization-type = "hvm"
  }
  most_recent = true
  owners      = ["099720109477"]
}

# source blocks are generated from your builders; a source can be referenced in
# build blocks. A build block runs provisioner and post-processors on a
# source. Read the documentation for source blocks here:
# https://www.packer.io/docs/from-1.5/blocks/source
source "amazon-ebs" "autogenerated_1" {
  access_key      = "${var.aws_access_key}"
  ami_description = "Ubuntu 16.04 LTS - expand root partition"
  ami_name        = "ubuntu-16-04-test-${local.timestamp}"
  encrypt_boot    = true
  launch_block_device_mappings {
    delete_on_termination = true
    device_name           = "/dev/sda1"
    volume_size           = 48
    volume_type           = "gp2"
  }
  region              = "${var.aws_region}"
  secret_key          = "${var.aws_secret_key}"
  spot_instance_types = ["t2.small", "t2.medium", "t2.large"]
  spot_price          = "0.0075"
  ssh_interface       = "session_manager"
  ssh_username        = "ubuntu"
  temporary_iam_instance_profile_policy_document {
    Statement {
      Action   = ["*"]
      Effect   = "Allow"
      Resource = ["*"]
    }
    Version = "2012-10-17"
  }
}

source "amazon-ebs" "named_builder" {
  access_key      = "${var.aws_access_key}"
  ami_description = "Ubuntu 16.04 LTS - expand root partition"
  ami_name        = "ubuntu-16-04-test-${local.timestamp}"
  encrypt_boot    = true
  launch_block_device_mappings {
    delete_on_termination = true
    device_name           = "/dev/sda1"
    volume_size           = 48
    volume_type           = "gp2"
  }
  region              = "${var.aws_region}"
  secret_key          = "${var.aws_secret_key}"
  spot_instance_types = ["t2.small", "t2.medium", "t2.large"]
  spot_price          = "0.0075"
  ssh_interface       = "session_manager"
  ssh_username        = "ubuntu"
  temporary_iam_instance_profile_policy_document {
    Statement {
      Action   = ["*"]
      Effect   = "Allow"
      Resource = ["*"]
    }
    Version = "2012-10-17"
  }
}

# a build block invokes sources and runs provisioning steps on them. The
# documentation for build blocks can be found here:
# https://www.packer.io/docs/from-1.5/blocks/build
build {
  sources = ["source.amazon-ebs.autogenerated_1", "source.amazon-ebs.named_builder"]

  provisioner "shell" {
    except      = ["amazon-ebs"]
    inline      = ["echo ${var.secret_account}", "echo ${build.ID}", "echo ${build.SSHPublicKey} | head -c 14", "echo ${path.root} is not ${path.cwd}", "echo ${packer.version}", "echo ${uuidv4()}"]
    max_retries = "5"
  }

  # template: hcl2_upgrade:2:38: executing "hcl2_upgrade" at <clean_resource_name>: error calling clean_resource_name: unhandled "clean_resource_name" call:
  # there is no way to automatically upgrade the "clean_resource_name" call.
  # Please manually upgrade to use custom validation rules, `replace(string, substring, replacement)` or `regex_replace(string, substring, replacement)`
  # Visit https://packer.io/docs/from-1.5/variables#custom-validation-rules , https://www.packer.io/docs/from-1.5/functions/string/replace or https://www.packer.io/docs/from-1.5/functions/string/regex_replace for more infos.
  provisioner "shell" {
    inline = ["echo mybuild-{{isotime | clean_resource_name}}"]
  }

  # template: hcl2_upgrade:2:35: executing "hcl2_upgrade" at <lower>: error calling lower: unhandled "lower" call:
  # there is no way to automatically upgrade the "lower" call.
  # Please manually upgrade to `lower(var.example)`
  # Visit https://www.packer.io/docs/from-1.5/functions/string/lower for more infos.
  provisioner "shell" {
    inline = ["echo {{ `SOMETHING` | lower }}"]
  }

  # template: hcl2_upgrade:2:35: executing "hcl2_upgrade" at <upper>: error calling upper: unhandled "upper" call:
  # there is no way to automatically upgrade the "upper" call.
  # Please manually upgrade to `upper(var.example)`
  # Visit https://www.packer.io/docs/from-1.5/functions/string/upper for more infos.
  provisioner "shell" {
    inline = ["echo {{ `something` | upper }}"]
  }

  # template: hcl2_upgrade:2:21: executing "hcl2_upgrade" at <split `some-string` `-` 0>: error calling split: unhandled "split" call:
  # there is no way to automatically upgrade the "split" call.
  # Please manually upgrade to `split(separator, string)`
  # Visit https://www.packer.io/docs/from-1.5/functions/string/split for more infos.
  provisioner "shell" {
    inline = ["echo {{ split `some-string` `-` 0 }}"]
  }

  # template: hcl2_upgrade:2:21: executing "hcl2_upgrade" at <replace_all `-` `/` build_name>: error calling replace_all: unhandled "replace_all" call:
  # there is no way to automatically upgrade the "replace_all" call.
  # Please manually upgrade to `replace(string, substring, replacement)` or `regex_replace(string, substring, replacement)`
  # Visit https://www.packer.io/docs/from-1.5/functions/string/replace or https://www.packer.io/docs/from-1.5/functions/string/regex_replace for more infos.
  provisioner "shell" {
    inline = ["echo {{ replace_all `-` `/` build_name }}"]
  }

  # template: hcl2_upgrade:2:21: executing "hcl2_upgrade" at <replace `some-string` `-` `/` 1>: error calling replace: unhandled "replace" call:
  # there is no way to automatically upgrade the "replace" call.
  # Please manually upgrade to `replace(string, substring, replacement)` or `regex_replace(string, substring, replacement)`
  # Visit https://www.packer.io/docs/from-1.5/functions/string/replace or https://www.packer.io/docs/from-1.5/functions/string/regex_replace for more infos.
  provisioner "shell" {
    inline = ["echo {{ replace `some-string` `-` `/` 1 }}"]
  }
  provisioner "shell-local" {
    inline  = ["sleep 100000"]
    only    = ["amazon-ebs"]
    timeout = "5s"
  }
  post-processor "amazon-import" {
    format         = "vmdk"
    license_type   = "BYOL"
    region         = "eu-west-3"
    s3_bucket_name = "hashicorp.adrien"
    tags = {
      Description = "packer amazon-import ${local.timestamp}"
    }
  }
  post-processors {
    post-processor "artifice" {
      keep_input_artifact = true
      files               = ["path/something.ova"]
      name                = "very_special_artifice_post-processor"
      only                = ["amazon-ebs"]
    }
    post-processor "amazon-import" {
      except         = ["amazon-ebs"]
      license_type   = "BYOL"
      s3_bucket_name = "hashicorp.adrien"
      tags = {
        Description = "packer amazon-import ${local.timestamp}"
      }
    }
  }
}
