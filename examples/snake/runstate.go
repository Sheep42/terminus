package main

import (
	"math/rand"
	"strconv"
	t "terminus"
	"time"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type RunState struct {
	*t.State
	scene       *CustomScene
	snake       []*t.Entity
	snakeLength int
	food        *t.Entity
	speed       float64
	dir         Direction
	elapsed     float64
	rand        *rand.Rand
	score       int
	scoreText   *t.Text
}

func NewRunState(scene *CustomScene) *RunState {

	s1 := rand.NewSource(time.Now().UnixNano())

	return &RunState{
		State: t.NewState(),
		scene: scene,
		rand:  rand.New(s1),
		score: 0,
	}

}

func (rs *RunState) OnEnter() {

	rs.snakeLength = 5
	rs.speed = 0.2

	rs.dir = Right
	rs.snake = []*t.Entity{}

	g := rs.scene.Game()
	gw, gh := g.ScreenSize()

	rs.food = t.NewSpriteEntity(rs.rand.Intn(gw), rs.rand.Intn(gh), 'o', t.Orange, t.Black)
	rs.scene.Add(rs.food)

	rs.scoreText = t.NewText(gw-15, 0, "Score: 0", t.White, t.Black)
	rs.scene.Add(rs.scoreText)

	for i := rs.snakeLength - 1; i >= 0; i-- {

		rs.snake = append(rs.snake, t.NewSpriteEntity(i+1, 5, 'o'))
		rs.scene.Add(rs.snake[rs.snakeLength-1-i])

	}

}

func (rs *RunState) OnExit() {

	for i := 0; i < rs.snakeLength; i++ {
		rs.scene.Remove(rs.snake[i])
	}

}

func (rs *RunState) Tick(delta float64) {

	// Update the scene - when the scene is updated
	// its child entities are also updated
	rs.scene.Scene.Update(delta)

	// update the timer
	rs.elapsed += delta

	// Check input
	g := rs.scene.Game()
	i := g.Input()

	if nil != i {

		if i.Key() == t.KeyRight && rs.dir != Left {
			rs.dir = Right
		} else if i.Key() == t.KeyLeft && rs.dir != Right {
			rs.dir = Left
		} else if i.Key() == t.KeyUp && rs.dir != Down {
			rs.dir = Up
		} else if i.Key() == t.KeyDown && rs.dir != Up {
			rs.dir = Down
		}

	}

	// only update the snake's position when the
	// timer exceeds the set speed
	if rs.elapsed < rs.speed {
		return
	}

	// reset the timer
	rs.elapsed = 0

	// Set the next position
	nextX, nextY := rs.snake[0].GetX(), rs.snake[0].GetY()

	if rs.dir == Right {
		nextX++
	} else if rs.dir == Left {
		nextX--
	} else if rs.dir == Up {
		nextY--
	} else if rs.dir == Down {
		nextY++
	}

	// check if the snake is off-screen
	gw, gh := g.ScreenSize()

	if nextX > gw || nextX < 0 ||
		nextY > gh || nextY < 0 {

		rs.scene.stateManager.ChangeState(rs.scene.endState)

	}

	// move the snake by moving the tail to the head
	rs.snake[rs.snakeLength-1].SetPosition(nextX, nextY)

	// then shift the rest of the snake in the slice
	tmp := rs.snake[rs.snakeLength-1]

	for i := rs.snakeLength - 1; i > 0; i-- {

		rs.snake[i] = rs.snake[i-1]

	}

	rs.snake[0] = tmp

	if rs.snake[0].Overlaps(rs.food) {

		newTailX, newTailY := rs.snake[rs.snakeLength-1].GetX(), rs.snake[rs.snakeLength-1].GetY()

		rs.snake = append(rs.snake, t.NewSpriteEntity(newTailX, newTailY, 'o'))
		rs.snakeLength++
		rs.scene.Add(rs.snake[rs.snakeLength-1])

		rs.food.SetPosition(rs.rand.Intn(gw), rs.rand.Intn(gh))
		rs.score += 5
		rs.scoreText.SetText("Score: " + strconv.Itoa(rs.score))

		if rs.speed > 0.05 && rs.score%25 == 0 {
			rs.speed -= 0.05
		}

	}

}
