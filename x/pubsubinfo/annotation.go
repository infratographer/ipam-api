package pubsubinfo

// AnnotationName is the value of the annotation when read during ent compilation
var AnnotationName = "INFRA9_PUBSUBHOOK"

// Annotation provides a ent.Annotaion spec
type Annotation struct {
	QueueName                string
	IsAdditionalSubjectField bool
}

// Name implements the ent Annotation interface.
func (a Annotation) Name() string {
	return AnnotationName
}

// AdditionalSubject marks this field as a field to return as an additional subject
func AdditionalSubject() *Annotation {
	return &Annotation{
		IsAdditionalSubjectField: true,
	}
}
