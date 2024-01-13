package model

import "fmt"

func (r ReproductiveMaterialType) ToPG() (string, error) {
	switch r {
	case ReproductiveMaterialTypeSeed:
		return "seed", nil
	case ReproductiveMaterialTypeCutting:
		return "cutting", nil
	case ReproductiveMaterialTypeGrafting:
		return "graft", nil
	default:
		return "", fmt.Errorf("unknown reproductive material type %s", r)
	}
}

func (r *ReproductiveMaterialType) FromPG(s string) error {
	switch s {
	case "seed":
		return r.UnmarshalGQL(ReproductiveMaterialTypeSeed)
	case "cutting":
		return r.UnmarshalGQL(ReproductiveMaterialTypeCutting)
	case "graft":
		return r.UnmarshalGQL(ReproductiveMaterialTypeGrafting)
	default:
		return fmt.Errorf("unknown reproductive material type %s", s)
	}
}

func (a AquisitionType) ToPG() (string, error) {
	switch a {
	case AquisitionTypeGathered:
		return "gathered", nil
	case AquisitionTypeGrown:
		return "grown", nil
	case AquisitionTypePurchased:
		return "bought", nil
	case AquisitionTypeHomeMade:
		return "home_made", nil
	default:
		return "", fmt.Errorf("unknown aquisition type %s", a)
	}
}

func (a *AquisitionType) FromPG(s string) error {
	switch s {
	case "gathered":
		return a.UnmarshalGQL(AquisitionTypeGathered)
	case "grown":
		return a.UnmarshalGQL(AquisitionTypeGrown)
	case "bought":
		return a.UnmarshalGQL(AquisitionTypePurchased)
	case "home_made":
		return a.UnmarshalGQL(AquisitionTypeHomeMade)
	default:
		return fmt.Errorf("unknown aquisition type %s", s)
	}
}

func (v Visibility) ToPG() (string, error) {
	switch v {
	case VisibilityInternal:
		return "internal", nil
	case VisibilityPublicFacing:
		return "public_facing", nil
	default:
		return "", fmt.Errorf("unknown visibility %s", v)
	}
}

func (v *Visibility) FromPG(s string) error {
	switch s {
	case "internal":
		return v.UnmarshalGQL(VisibilityInternal)
	case "public_facing":
		return v.UnmarshalGQL(VisibilityPublicFacing)
	default:
		return fmt.Errorf("unknown visibility %s", s)
	}
}
