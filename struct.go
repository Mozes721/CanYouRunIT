package main

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

type allGames struct {
	Applist struct {
		Apps []struct {
			Appid int    `json:"appid"`
			Name  string `json:"name"`
		} `json:"apps"`
	} `json:"applist"`
}

type PcRequirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}
type MacRequirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}
type PriceOverview struct {
	Currency         string `json:"currency"`
	Initial          int    `json:"initial"`
	Final            int    `json:"final"`
	DiscountPercent  int    `json:"discount_percent"`
	InitialFormatted string `json:"initial_formatted"`
	FinalFormatted   string `json:"final_formatted"`
}
type Subs struct {
	Packageid                int    `json:"packageid"`
	PercentSavingsText       string `json:"percent_savings_text"`
	PercentSavings           int    `json:"percent_savings"`
	OptionText               string `json:"option_text"`
	OptionDescription        string `json:"option_description"`
	CanGetFreeLicense        string `json:"can_get_free_license"`
	IsFreeLicense            bool   `json:"is_free_license"`
	PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
}
type PackageGroups struct {
	Name                    string `json:"name"`
	Title                   string `json:"title"`
	Description             string `json:"description"`
	SelectionText           string `json:"selection_text"`
	SaveText                string `json:"save_text"`
	DisplayType             int    `json:"display_type"`
	IsRecurringSubscription string `json:"is_recurring_subscription"`
	Subs                    []Subs `json:"subs"`
}
type Platforms struct {
	Windows bool `json:"windows"`
	Mac     bool `json:"mac"`
	Linux   bool `json:"linux"`
}
type Categories struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
type Genres struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
type Screenshots struct {
	ID            int    `json:"id"`
	PathThumbnail string `json:"path_thumbnail"`
	PathFull      string `json:"path_full"`
}
type Webm struct {
	Num480 string `json:"480"`
	Max    string `json:"max"`
}
type Mp4 struct {
	Num480 string `json:"480"`
	Max    string `json:"max"`
}
type Movies struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Webm      Webm   `json:"webm"`
	Mp4       Mp4    `json:"mp4"`
	Highlight bool   `json:"highlight"`
}
type Recommendations struct {
	Total int `json:"total"`
}
type Highlighted struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
type Achievements struct {
	Total       int           `json:"total"`
	Highlighted []Highlighted `json:"highlighted"`
}
type ReleaseDate struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}
type SupportInfo struct {
	URL   string `json:"url"`
	Email string `json:"email"`
}
type ContentDescriptors struct {
	Ids   []interface{} `json:"ids"`
	Notes interface{}   `json:"notes"`
}
type Data struct {
	PcRequirements PcRequirements `json:"pc_requirements"`
	// MacRequirements   MacRequirements `json:"mac_requirements"`
	// LinuxRequirements []interface{}   `json:"linux_requirements"`
}
type GameNum struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type GameData struct {
	SoundCard string
	Storage   string
	Network   string
	Graphics  string
	Memory    string
	Processor string
}
