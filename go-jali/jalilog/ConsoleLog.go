package jalilog

type ConsoleLog struct {
}

func NewConsoleLog(config *ConsoleLogConfig) (*ConsoleLog, error) {
	//	newConfig, err := buildConsoleConfig(config, DefaultConsoleLogConfig())
	//
	//	if err != nil {
	//		return _, err
	//	}

	return &ConsoleLog{}
}

func (log *ConsoleLog) Log(info *LogInformer) {

}

func (log *ConsoleLog) Verbose(message string) {

}

func buildConsoleConfig(config *ConsoleLogConfig, defaultConfig *ConsoleLogConfig) (*ConsoleLog, error) {
	//	var newConfig ConsoleLogConfig
	//	if config != nil {
	//		newConfig = config
	//	} else {
	//		newConfig = defaultConfig
	//	}

	return &ConsoleLog{}, _
}
