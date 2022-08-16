type DXE_menu_01 struct {
	Version    string `json:"version"`
	Categories []struct {
		Tag          string `json:"tag"`
		Name         string `json:"name"`
		DesktopImage string `json:"desktopImage,omitempty"`
		MobileImage  string `json:"mobileImage"`
		Items        []struct {
			Tag                   string   `json:"tag"`
			Name                  string   `json:"name"`
			Description           string   `json:"description"`
			ItemType              string   `json:"itemType"`
			ItemClass             string   `json:"itemClass"`
			ItemGroupType         string   `json:"itemGroupType"`
			BeverageTypeGrouping  string   `json:"beverageTypeGrouping"`
			DayPart               string   `json:"dayPart"`
			DesktopImage          string   `json:"desktopImage"`
			DotComVisible         bool     `json:"dotComVisible"`
			Meal                  bool     `json:"meal"`
			MobileImage           string   `json:"mobileImage"`
			PdpImages             string   `json:"pdpImages"`
			SaltLaw               string   `json:"saltLaw"`
			Status                string   `json:"status"`
			Disabled              bool     `json:"disabled"`
			ItemGroupID           int      `json:"itemGroupId"`
			ThirdPartyDescription string   `json:"thirdPartyDescription"`
			HoldTime              int      `json:"holdTime"`
			ItemNickname          string   `json:"itemNickname"`
			EarlyAccessTiers      []string `json:"earlyAccessTiers"`
			ObjectType            string   `json:"objectType"`
			IsEdible              bool     `json:"isEdible"`
			IsSellable            bool     `json:"isSellable"`
		} `json:"items"`
		Description string `json:"description,omitempty"`
	} `json:"categories"`
	ItemGroups []struct {
		ItemGroupID int `json:"itemGroupId"`
		Items       []struct {
			Tag                  string `json:"tag"`
			Name                 string `json:"name"`
			ItemType             string `json:"itemType"`
			ItemClass            string `json:"itemClass"`
			ItemGroupType        string `json:"itemGroupType"`
			BeverageTypeGrouping string `json:"beverageTypeGrouping"`
			DayPart              string `json:"dayPart"`
			DesktopImage         string `json:"desktopImage"`
			DotComVisible        bool   `json:"dotComVisible"`
			Meal                 bool   `json:"meal"`
			MobileImage          string `json:"mobileImage"`
			PdpImages            string `json:"pdpImages"`
			SaltLaw              string `json:"saltLaw"`
			Status               string `json:"status"`
			ModifierType         string `json:"modifierType"`
			HoldTime             int    `json:"holdTime"`
			ObjectType           string `json:"objectType"`
			IsEdible             bool   `json:"isEdible"`
			IsSellable           bool   `json:"isSellable"`
		} `json:"items"`
	} `json:"itemGroups"`
}


