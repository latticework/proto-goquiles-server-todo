package jaliconfig

type CompositeConfig struct {
	Configurators []Configurator
}

func DefaultCompositeConfig() *CompositeConfig {

	jsonConfig, _ := NewJsonConfiguration(&DefaultJsonConfig)

	configurators := []Configurator{
		jsonConfig,
		// TODO: Complete implementation of CompositeConfiguration
		//		NewEnvirionmentConfiguration(DefaultEnvironmentConfig),
		//		NewCommandLineConfiguration(DefaultCommandLineConfig),
	}

	return CompositeConfig{Configurators: configurators}
}
