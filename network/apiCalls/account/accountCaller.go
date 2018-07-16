package account

type AccountCaller interface{
	GetAchievements() (error)
}

type accountCaller struct {

}
