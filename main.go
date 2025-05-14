package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	go startWebServer()
	fmt.Println("🟢 Sistem Kaynak İzleyici Başladı")
	logToFile("🟢 Sistem Kaynak İzleyici Başladı")

	for {
		now := time.Now().Format("2006-01-02 15:04:05")
		logLine := fmt.Sprintf("⌛ Zaman: %s", now)
		fmt.Println(logLine)
		logToFile(logLine)

		printAndLogCPU()
		printAndLogRAM()
		printAndLogDisk()

		fmt.Println("------------------------------")
		logToFile("------------------------------")

		time.Sleep(10 * time.Second)
	}
}

func printAndLogCPU() {
	percent, _ := cpu.Percent(0, false)
	cpuUsage := percent[0]
	line := fmt.Sprintf("🧠 CPU Kullanımı: %.2f%%", cpuUsage)
	fmt.Println(line)
	logToFile(line)

	if cpuUsage > 80 {
		warning := "⚠️  UYARI: CPU kullanımı yüksek!"
		fmt.Println(warning)
		logToFile(warning)
		sendMail("⚠️ CPU Uyarısı", fmt.Sprintf("CPU kullanımı çok yüksek: %.2f%%", cpuUsage))
	}
}

func printAndLogRAM() {
	vmStat, _ := mem.VirtualMemory()
	ramUsage := vmStat.UsedPercent
	line := fmt.Sprintf("💾 RAM Kullanımı: %.2f%% (%v / %v)", ramUsage, byteToGB(vmStat.Used), byteToGB(vmStat.Total))
	fmt.Println(line)
	logToFile(line)

	if ramUsage > 85 {
		warning := "⚠️  UYARI: RAM kullanımı kritik seviyede!"
		fmt.Println(warning)
		logToFile(warning)
		sendMail("⚠️ RAM Uyarısı", fmt.Sprintf("RAM kullanımı çok yüksek: %.2f%%", ramUsage))
	}
}

func printAndLogDisk() {
	diskStat, _ := disk.Usage("/")
	diskUsage := diskStat.UsedPercent
	line := fmt.Sprintf("🗂️  Disk Kullanımı: %.2f%% (%v / %v)", diskUsage, byteToGB(diskStat.Used), byteToGB(diskStat.Total))
	fmt.Println(line)
	logToFile(line)

	if diskUsage > 90 {
		warning := "⚠️  UYARI: Disk kullanımı çok yüksek!"
		fmt.Println(warning)
		logToFile(warning)
		sendMail("⚠️ Disk Uyarısı", fmt.Sprintf("Disk kullanımı çok yüksek: %.2f%%", diskUsage))
	}
}

func logToFile(logLine string) {
	logDir := "logs"
	os.MkdirAll(logDir, os.ModePerm)

	date := time.Now().Format("2006-01-02")
	logFile := filepath.Join(logDir, fmt.Sprintf("%s.log", date))

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Log dosyası açılamadı:", err)
		return
	}
	defer f.Close()

	timestamp := time.Now().Format("15:04:05")
	line := fmt.Sprintf("[%s] %s\n", timestamp, logLine)
	f.WriteString(line)
}

func byteToGB(b uint64) string {
	return fmt.Sprintf("%.2f GB", float64(b)/1024/1024/1024)
}
