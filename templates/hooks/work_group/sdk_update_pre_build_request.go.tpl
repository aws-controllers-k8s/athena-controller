	if delta.DifferentAt("Spec.Tags") {
		if err = rm.syncTags(ctx, desired, latest); err != nil {
			return nil, err
		}
	}
	if !delta.DifferentExcept("Spec.Tags") {
		// Tags are synced above; UpdateWorkGroup has nothing to send.
		return desired, nil
	}
