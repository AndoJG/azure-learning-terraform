# azure-learning-terraform

Azure Terraform modules designed and written for use with Terragrunt

For usage examples, see:  
[azure-learning-terragrunt](https://github.com/breadwatcher/azure-learning-terragrunt)

## Requirements

* [Terraform](https://www.terraform.io/) - Developed with v 0.14.8
* [Terragrunt](https://terragrunt.gruntwork.io/) - Developed with v 0.28.12
* [Go](https://golang.org/) - Developed with v 1.16.2

All of these things can be installed via Homebrew

## Modules

Modules in this project are designed and written with the expressed intent of being used with Terragrunt.

### What does that mean for writing a module?

* Modules should not define a provider, as that can and should be done with Terragrunt
* Any parameter value in the module that might change should be configurable as an input variable

## Testing

Module testing is done with [Terratest](https://terratest.gruntwork.io/), a Golang package that provides patterns and functions to help test infrastructure
