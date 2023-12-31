// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/99designs/gqlgen/plugin/federation/fedruntime"
)

var (
	ErrUnknownType  = errors.New("unknown type")
	ErrTypeNotFound = errors.New("type not found")
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.Service{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) []fedruntime.Entity {
	list := make([]fedruntime.Entity, len(representations))

	repsMap := map[string]struct {
		i []int
		r []map[string]interface{}
	}{}

	// We group entities by typename so that we can parallelize their resolution.
	// This is particularly helpful when there are entity groups in multi mode.
	buildRepresentationGroups := func(reps []map[string]interface{}) {
		for i, rep := range reps {
			typeName, ok := rep["__typename"].(string)
			if !ok {
				// If there is no __typename, we just skip the representation;
				// we just won't be resolving these unknown types.
				ec.Error(ctx, errors.New("__typename must be an existing string"))
				continue
			}

			_r := repsMap[typeName]
			_r.i = append(_r.i, i)
			_r.r = append(_r.r, rep)
			repsMap[typeName] = _r
		}
	}

	isMulti := func(typeName string) bool {
		switch typeName {
		default:
			return false
		}
	}

	resolveEntity := func(ctx context.Context, typeName string, rep map[string]interface{}, idx []int, i int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {
		case "Attendance":
			resolverName, err := entityResolverNameForAttendance(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Attendance": %w`, err)
			}
			switch resolverName {

			case "findAttendanceByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findAttendanceByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindAttendanceByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Attendance": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Board":
			resolverName, err := entityResolverNameForBoard(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Board": %w`, err)
			}
			switch resolverName {

			case "findBoardByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findBoardByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindBoardByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Board": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Campaign":
			resolverName, err := entityResolverNameForCampaign(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Campaign": %w`, err)
			}
			switch resolverName {

			case "findCampaignByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findCampaignByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindCampaignByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Campaign": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Collaborator":
			resolverName, err := entityResolverNameForCollaborator(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Collaborator": %w`, err)
			}
			switch resolverName {

			case "findCollaboratorByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findCollaboratorByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindCollaboratorByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Collaborator": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Comment":
			resolverName, err := entityResolverNameForComment(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Comment": %w`, err)
			}
			switch resolverName {

			case "findCommentByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findCommentByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindCommentByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Comment": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Company":
			resolverName, err := entityResolverNameForCompany(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Company": %w`, err)
			}
			switch resolverName {

			case "findCompanyByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findCompanyByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindCompanyByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Company": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Contact":
			resolverName, err := entityResolverNameForContact(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Contact": %w`, err)
			}
			switch resolverName {

			case "findContactByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findContactByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindContactByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Contact": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Coupon":
			resolverName, err := entityResolverNameForCoupon(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Coupon": %w`, err)
			}
			switch resolverName {

			case "findCouponByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findCouponByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindCouponByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Coupon": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Expense":
			resolverName, err := entityResolverNameForExpense(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Expense": %w`, err)
			}
			switch resolverName {

			case "findExpenseByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findExpenseByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindExpenseByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Expense": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "File":
			resolverName, err := entityResolverNameForFile(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "File": %w`, err)
			}
			switch resolverName {

			case "findFileByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findFileByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindFileByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "File": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Goal":
			resolverName, err := entityResolverNameForGoal(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Goal": %w`, err)
			}
			switch resolverName {

			case "findGoalByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findGoalByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindGoalByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Goal": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Invoice":
			resolverName, err := entityResolverNameForInvoice(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Invoice": %w`, err)
			}
			switch resolverName {

			case "findInvoiceByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findInvoiceByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindInvoiceByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Invoice": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Link":
			resolverName, err := entityResolverNameForLink(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Link": %w`, err)
			}
			switch resolverName {

			case "findLinkByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findLinkByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindLinkByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Link": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "List":
			resolverName, err := entityResolverNameForList(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "List": %w`, err)
			}
			switch resolverName {

			case "findListByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findListByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindListByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "List": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Member":
			resolverName, err := entityResolverNameForMember(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Member": %w`, err)
			}
			switch resolverName {

			case "findMemberByID":
				id0, err := ec.unmarshalNID2string(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findMemberByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindMemberByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Member": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Membership":
			resolverName, err := entityResolverNameForMembership(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Membership": %w`, err)
			}
			switch resolverName {

			case "findMembershipByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findMembershipByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindMembershipByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Membership": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Order":
			resolverName, err := entityResolverNameForOrder(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Order": %w`, err)
			}
			switch resolverName {

			case "findOrderByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findOrderByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindOrderByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Order": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Organization":
			resolverName, err := entityResolverNameForOrganization(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Organization": %w`, err)
			}
			switch resolverName {

			case "findOrganizationByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findOrganizationByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindOrganizationByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Organization": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Portfolio":
			resolverName, err := entityResolverNameForPortfolio(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Portfolio": %w`, err)
			}
			switch resolverName {

			case "findPortfolioByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findPortfolioByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindPortfolioByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Portfolio": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Post":
			resolverName, err := entityResolverNameForPost(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Post": %w`, err)
			}
			switch resolverName {

			case "findPostByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findPostByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindPostByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Post": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Quote":
			resolverName, err := entityResolverNameForQuote(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Quote": %w`, err)
			}
			switch resolverName {

			case "findQuoteByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findQuoteByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindQuoteByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Quote": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Reaction":
			resolverName, err := entityResolverNameForReaction(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Reaction": %w`, err)
			}
			switch resolverName {

			case "findReactionByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findReactionByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindReactionByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Reaction": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Share":
			resolverName, err := entityResolverNameForShare(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Share": %w`, err)
			}
			switch resolverName {

			case "findShareByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findShareByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindShareByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Share": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Task":
			resolverName, err := entityResolverNameForTask(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Task": %w`, err)
			}
			switch resolverName {

			case "findTaskByUIDAndSharesAndAssignee":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findTaskByUIDAndSharesAndAssignee(): %w`, err)
				}
				id1, err := ec.unmarshalOID2ᚕstringᚄ(ctx, rep["shares"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 1 for findTaskByUIDAndSharesAndAssignee(): %w`, err)
				}
				id2, err := ec.unmarshalNID2string(ctx, rep["assignee"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 2 for findTaskByUIDAndSharesAndAssignee(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindTaskByUIDAndSharesAndAssignee(ctx, id0, id1, id2)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Task": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "User":
			resolverName, err := entityResolverNameForUser(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "User": %w`, err)
			}
			switch resolverName {

			case "findUserByID":
				id0, err := ec.unmarshalNID2string(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findUserByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindUserByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "User": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Wishlist":
			resolverName, err := entityResolverNameForWishlist(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Wishlist": %w`, err)
			}
			switch resolverName {

			case "findWishlistByUID":
				id0, err := ec.unmarshalNID2string(ctx, rep["uid"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findWishlistByUID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindWishlistByUID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Wishlist": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}

		}
		return fmt.Errorf("%w: %s", ErrUnknownType, typeName)
	}

	resolveManyEntities := func(ctx context.Context, typeName string, reps []map[string]interface{}, idx []int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {

		default:
			return errors.New("unknown type: " + typeName)
		}
	}

	resolveEntityGroup := func(typeName string, reps []map[string]interface{}, idx []int) {
		if isMulti(typeName) {
			err := resolveManyEntities(ctx, typeName, reps, idx)
			if err != nil {
				ec.Error(ctx, err)
			}
		} else {
			// if there are multiple entities to resolve, parallelize (similar to
			// graphql.FieldSet.Dispatch)
			var e sync.WaitGroup
			e.Add(len(reps))
			for i, rep := range reps {
				i, rep := i, rep
				go func(i int, rep map[string]interface{}) {
					err := resolveEntity(ctx, typeName, rep, idx, i)
					if err != nil {
						ec.Error(ctx, err)
					}
					e.Done()
				}(i, rep)
			}
			e.Wait()
		}
	}
	buildRepresentationGroups(representations)

	switch len(repsMap) {
	case 0:
		return list
	case 1:
		for typeName, reps := range repsMap {
			resolveEntityGroup(typeName, reps.r, reps.i)
		}
		return list
	default:
		var g sync.WaitGroup
		g.Add(len(repsMap))
		for typeName, reps := range repsMap {
			go func(typeName string, reps []map[string]interface{}, idx []int) {
				resolveEntityGroup(typeName, reps, idx)
				g.Done()
			}(typeName, reps.r, reps.i)
		}
		g.Wait()
		return list
	}
}

func entityResolverNameForAttendance(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findAttendanceByUID", nil
	}
	return "", fmt.Errorf("%w for Attendance", ErrTypeNotFound)
}

func entityResolverNameForBoard(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findBoardByUID", nil
	}
	return "", fmt.Errorf("%w for Board", ErrTypeNotFound)
}

func entityResolverNameForCampaign(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findCampaignByUID", nil
	}
	return "", fmt.Errorf("%w for Campaign", ErrTypeNotFound)
}

func entityResolverNameForCollaborator(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findCollaboratorByUID", nil
	}
	return "", fmt.Errorf("%w for Collaborator", ErrTypeNotFound)
}

func entityResolverNameForComment(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findCommentByUID", nil
	}
	return "", fmt.Errorf("%w for Comment", ErrTypeNotFound)
}

func entityResolverNameForCompany(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findCompanyByUID", nil
	}
	return "", fmt.Errorf("%w for Company", ErrTypeNotFound)
}

func entityResolverNameForContact(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findContactByUID", nil
	}
	return "", fmt.Errorf("%w for Contact", ErrTypeNotFound)
}

func entityResolverNameForCoupon(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findCouponByUID", nil
	}
	return "", fmt.Errorf("%w for Coupon", ErrTypeNotFound)
}

func entityResolverNameForExpense(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findExpenseByUID", nil
	}
	return "", fmt.Errorf("%w for Expense", ErrTypeNotFound)
}

func entityResolverNameForFile(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findFileByUID", nil
	}
	return "", fmt.Errorf("%w for File", ErrTypeNotFound)
}

func entityResolverNameForGoal(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findGoalByUID", nil
	}
	return "", fmt.Errorf("%w for Goal", ErrTypeNotFound)
}

func entityResolverNameForInvoice(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findInvoiceByUID", nil
	}
	return "", fmt.Errorf("%w for Invoice", ErrTypeNotFound)
}

func entityResolverNameForLink(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findLinkByUID", nil
	}
	return "", fmt.Errorf("%w for Link", ErrTypeNotFound)
}

func entityResolverNameForList(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findListByUID", nil
	}
	return "", fmt.Errorf("%w for List", ErrTypeNotFound)
}

func entityResolverNameForMember(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findMemberByID", nil
	}
	return "", fmt.Errorf("%w for Member", ErrTypeNotFound)
}

func entityResolverNameForMembership(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findMembershipByUID", nil
	}
	return "", fmt.Errorf("%w for Membership", ErrTypeNotFound)
}

func entityResolverNameForOrder(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findOrderByUID", nil
	}
	return "", fmt.Errorf("%w for Order", ErrTypeNotFound)
}

func entityResolverNameForOrganization(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findOrganizationByUID", nil
	}
	return "", fmt.Errorf("%w for Organization", ErrTypeNotFound)
}

func entityResolverNameForPortfolio(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findPortfolioByUID", nil
	}
	return "", fmt.Errorf("%w for Portfolio", ErrTypeNotFound)
}

func entityResolverNameForPost(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findPostByUID", nil
	}
	return "", fmt.Errorf("%w for Post", ErrTypeNotFound)
}

func entityResolverNameForQuote(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findQuoteByUID", nil
	}
	return "", fmt.Errorf("%w for Quote", ErrTypeNotFound)
}

func entityResolverNameForReaction(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findReactionByUID", nil
	}
	return "", fmt.Errorf("%w for Reaction", ErrTypeNotFound)
}

func entityResolverNameForShare(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findShareByUID", nil
	}
	return "", fmt.Errorf("%w for Share", ErrTypeNotFound)
}

func entityResolverNameForTask(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		m = rep
		if _, ok = m["shares"]; !ok {
			break
		}
		m = rep
		if _, ok = m["assignee"]; !ok {
			break
		}
		return "findTaskByUIDAndSharesAndAssignee", nil
	}
	return "", fmt.Errorf("%w for Task", ErrTypeNotFound)
}

func entityResolverNameForUser(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findUserByID", nil
	}
	return "", fmt.Errorf("%w for User", ErrTypeNotFound)
}

func entityResolverNameForWishlist(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["uid"]; !ok {
			break
		}
		return "findWishlistByUID", nil
	}
	return "", fmt.Errorf("%w for Wishlist", ErrTypeNotFound)
}
