	// CreateDataCatalog returns a fully populated DataCatalog body only for the
	// FEDERATED catalog type. For LAMBDA, HIVE, and GLUE catalogs the response
	// does not reliably echo the Spec fields, so the generated create-output
	// mapping can overwrite user-provided values (including the primary key
	// Name) with nil. Preserve the desired Spec here; the readOne (sdkFind)
	// path reads the authoritative state from GetDataCatalog on the next
	// reconcile.
	ko.Spec.Name = desired.ko.Spec.Name
	ko.Spec.Type = desired.ko.Spec.Type
	ko.Spec.Description = desired.ko.Spec.Description
	ko.Spec.Parameters = desired.ko.Spec.Parameters
