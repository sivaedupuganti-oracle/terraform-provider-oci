
# oci_database_db_home_patches

## DbHomePatch DataSource

Gets a list of db_home_patches.

### List Operation
Lists patches applicable to the requested database home.

The following arguments are supported:

* `db_home_id` - (Required) The database home [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


The following attributes are exported:

* `patches` - The list of patches.

### Example Usage

```hcl
data "oci_database_db_home_patches" "test_db_home_patches" {
	#Required
	db_home_id = "${oci_database_db_home.test_db_home.id}"
}
```
### DbHomePatch Reference

The following attributes are exported:

* `available_actions` - Actions that can possibly be performed using this patch.
* `description` - The text describing this patch package.
* `id` - The OCID of the patch.
* `last_action` - Action that is currently being performed or was completed last.
* `lifecycle_details` - A descriptive text associated with the lifecycleState. Typically can contain additional displayable text. 
* `state` - The current state of the patch as a result of lastAction.
* `time_released` - The date and time that the patch was released.
* `version` - The version of this patch package.
