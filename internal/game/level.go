package game

type LevelImpl struct {
	sections       []Section
	currentSection int
}

func NewLevel() *LevelImpl {
	return &LevelImpl{
		sections:       make([]Section, 0),
		currentSection: 0,
	}
}

func (l *LevelImpl) AddSection(section Section) {
	l.sections = append(l.sections, section)
}

func (l *LevelImpl) CurrentSection() Section {
	return l.sections[l.currentSection]
}

func (l *LevelImpl) Init() {
	currentSection := l.sections[l.currentSection]
	currentSection.SetPlayerPositionAtLeft()
}
