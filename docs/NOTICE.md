### What has changed?
This module has been scoped down to only handle inventory 
files and manage dedicated host inside those files. The pre-existing module
`go-ansible-db` is a great outline for vault management/inventory management
should a user need both. The focus of these changes are to 
update/iterate over outdated code and to make changes to how inventory types
are managed. With support for yaml style inventory files this module `goggle`
is a shared resource for parsing ini/yaml formatted inventory files and
grabbing host information for such file types.

### What will be added?
- Inventory file type support for yaml
- Ansible specific customization for plugin configuration
- Support for ansible-core and its current libraries
