package core

type Rule func() FieldError

func Execute(
	rules []Rule,
) error {

	var errs Errors

	for _, rule := range rules {

		if err := rule(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
