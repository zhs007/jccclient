package jccclient

// VERSION - jccclient version
const VERSION = "0.1.45"

// MaxMultiFailTimes - If the client fail {{MaxMultiFailTimes}} times in succession, it need to ignore some task
const MaxMultiFailTimes = 3

// EffectiveFailTime - If the interval between 2 failures is less than {{EffectiveFailTime}} seconds, it should be counted as a multiple failure.
const EffectiveFailTime = 3 * 60

// IgnoreTaskTime - The client will ignore the task for {{IgnoreTaskTime}} seconds.
const IgnoreTaskTime = 5 * 60

// DefaultSleepTime - After the finished a task, rest for {{DefaultSleepTime}} seconds.
const DefaultSleepTime = 5

// DBName - jccclient database name
const DBName = "jccclient"
