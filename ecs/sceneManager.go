package ecs

type SceneManager interface {
	Tick(tickCnt uint64)
	ChangeScene(scene Scene)
}

func NewSceneManager() SceneManager {
	return &implSceneManager{}
}

type implSceneManager struct {
	scene Scene
}

func (s *implSceneManager) Tick(tickCnt uint64) {
	if s.scene != nil {
		s.scene.Tick(tickCnt)
	}
}

func (s *implSceneManager) ChangeScene(scene Scene) {
	if s.scene != nil {
		s.scene.OnDestroy()
	}
	s.scene = scene
	s.scene.Init()
}
