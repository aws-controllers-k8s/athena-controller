	// Athena injects server-managed parameters into data catalogs that the user
	// never specified (for example "catalog" for LAMBDA/HIVE and "sdk-version"
	// for HIVE). Left in Spec.Parameters they produce a permanent diff that
	// drives the controller to call UpdateDataCatalog on every reconcile, which
	// in turn drops the reconstructed ARN from status. Keep only the parameters
	// the user actually declared in the desired spec.
	if ko.Spec.Parameters != nil && r.ko.Spec.Parameters != nil {
		for k := range ko.Spec.Parameters {
			if _, userSet := r.ko.Spec.Parameters[k]; !userSet {
				delete(ko.Spec.Parameters, k)
			}
		}
	}

	if ko.Status.ACKResourceMetadata != nil {
		// We need to build the resourceARN from accountID and region,
		// since it is not directly returned by the API.
		resourceARN := ackv1alpha1.AWSResourceName(fmt.Sprintf("arn:aws:athena:%s:%s:datacatalog/%s",
		*ko.Status.ACKResourceMetadata.Region, *ko.Status.ACKResourceMetadata.OwnerAccountID, *ko.Spec.Name))

		// Set resourceARN to status
		ko.Status.ACKResourceMetadata.ARN = &resourceARN

		// Now we can fetch the tags using the manually constructed ARN
		tags, err := rm.getTags(ctx, string(resourceARN))
		if err != nil {
			return nil, err
		}
		ko.Spec.Tags = tags
	}
