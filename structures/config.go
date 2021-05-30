package structures

type Configuration struct {
	HOST     string   `yaml:"host"`
	PORT     string   `yaml:"port"`
	PASSWORD string   `yaml:"password"`
	BUTTONS  []Button `yaml:"buttons,flow"`
}

type Button struct {
	Name         string   `yaml:"name"`
	DisplayName  string   `yaml:"display_name"`
	ShellCommand string   `yaml:"shell_command"`
	Keys         []string `yaml:"keys"`
	Logo         string   `yaml:"logo"`
	HasShift     bool     `yaml:"shift"`
	HasRShift    bool     `yaml:"rshift"`
	HasCtrl      bool     `yaml:"ctrl"`
	HasRCtrl     bool     `yaml:"rctrl"`
	HasAlt       bool     `yaml:"alt"`
	HasRAlt      bool     `yaml:"ralt"`
}

type MinifiedButton struct {
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	DisplayName string `json:"display_name"`
}

type ClickButtonRequest struct {
	Password    string `json:"password"`
	CommandName string `json:"command"`
}

type PasswordRequest struct {
	Password string `json:"password"`
}
