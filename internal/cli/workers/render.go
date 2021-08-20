package workers

import (
	"io"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spiral/roadrunner/v2/pkg/state/job"
	"github.com/spiral/roadrunner/v2/pkg/state/process"
)

// WorkerTable renders table with information about rr server workers.
func WorkerTable(writer io.Writer, workers []*process.State) *tablewriter.Table {
	tw := tablewriter.NewWriter(writer)
	tw.SetHeader([]string{"PID", "Status", "Execs", "Memory", "CPU%", "Created"})
	tw.SetColMinWidth(0, 7)
	tw.SetColMinWidth(1, 9)
	tw.SetColMinWidth(2, 7)
	tw.SetColMinWidth(3, 7)
	tw.SetColMinWidth(4, 7)
	tw.SetColMinWidth(5, 18)

	for i := 0; i < len(workers); i++ {
		tw.Append([]string{
			strconv.Itoa(workers[i].Pid),
			renderStatus(workers[i].Status),
			renderJobs(workers[i].NumJobs),
			humanize.Bytes(workers[i].MemoryUsage),
			renderCPU(workers[i].CPUPercent),
			renderAlive(time.Unix(0, workers[i].Created)),
		})
	}

	return tw
}

// ServiceWorkerTable renders table with information about rr server workers.
func ServiceWorkerTable(writer io.Writer, workers []*process.State) *tablewriter.Table {
	tw := tablewriter.NewWriter(writer)
	tw.SetAutoWrapText(false)
	tw.SetHeader([]string{"PID", "Memory", "CPU%", "Command"})
	tw.SetColMinWidth(0, 7)
	tw.SetColMinWidth(1, 7)
	tw.SetColMinWidth(2, 7)
	tw.SetColMinWidth(3, 18)
	tw.SetAlignment(tablewriter.ALIGN_LEFT)

	for i := 0; i < len(workers); i++ {
		tw.Append([]string{
			strconv.Itoa(workers[i].Pid),
			humanize.Bytes(workers[i].MemoryUsage),
			renderCPU(workers[i].CPUPercent),
			workers[i].Command,
		})
	}

	return tw
}

// JobsTable renders table with information about rr server jobs.
func JobsTable(writer io.Writer, jobs []*job.State) *tablewriter.Table {
	tw := tablewriter.NewWriter(writer)
	tw.SetAutoWrapText(false)
	tw.SetHeader([]string{"Pipeline", "Driver", "Queue", "Active", "Delayed", "Reserved"})
	tw.SetColMinWidth(0, 10)
	tw.SetColMinWidth(1, 10)
	tw.SetColMinWidth(2, 7)
	tw.SetColMinWidth(3, 15)
	tw.SetColMinWidth(4, 10)
	tw.SetColMinWidth(5, 10)
	tw.SetColMinWidth(6, 10)
	tw.SetAlignment(tablewriter.ALIGN_LEFT)

	for i := 0; i < len(jobs); i++ {
		tw.Append([]string{
			renderReady(jobs[i].Ready),
			jobs[i].Pipeline,
			jobs[i].Driver,
			jobs[i].Queue,
			strconv.Itoa(int(jobs[i].Active)),
			strconv.Itoa(int(jobs[i].Delayed)),
			strconv.Itoa(int(jobs[i].Reserved)),
		})
	}

	return tw
}

func renderReady(ready bool) string {
	if ready == true {
		return "READY"
	}
	return "PAUSED/STOPPED"
}

//go:inline
func renderCPU(cpu float64) string {
	return strconv.FormatFloat(cpu, 'f', 2, 64)
}

func renderStatus(status string) string {
	switch status {
	case "inactive":
		return color.YellowString("inactive")
	case "ready":
		return color.CyanString("ready")
	case "working":
		return color.GreenString("working")
	case "invalid":
		return color.YellowString("invalid")
	case "stopped":
		return color.RedString("stopped")
	case "errored":
		return color.RedString("errored")
	default:
		return status
	}
}

func renderJobs(number uint64) string {
	return humanize.Comma(int64(number))
}

func renderAlive(t time.Time) string {
	return humanize.RelTime(t, time.Now(), "ago", "")
}
