package task


type Task struct {
  Src string
  DestScheme string
  DestBkt string
  DestKey string
}

type FeedBack struct {
  Task *Task
  Hash uint32 // defaults 0
  HashZ uint32 // defaults 0
  Err error
}
