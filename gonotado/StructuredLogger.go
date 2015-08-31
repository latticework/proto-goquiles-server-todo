package notado

type StructuredLogger interface {
	Log(info *LogInformer)
	Verbose(message string)
}
