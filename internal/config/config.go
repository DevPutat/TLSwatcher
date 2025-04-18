package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/DevPutat/TLSwatcher/internal/types"
)

func parse(reader io.Reader) ([]types.Domain, error) {
	res := []types.Domain{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if len(url) == 0 {
			continue
		}
		res = append(res, types.Domain{
			Url:         url,
			IsConnected: false,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func Domains(confPath string) ([]types.Domain, error) {
	if _, err := os.Stat(confPath); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	file, err := os.Open(confPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return parse(file)
}

func InputDomains(confPath string) error {
	file, err := os.Create(confPath)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Println("Введите список доменов, разделяя их пробелами")
	err = write(os.Stdin, file)
	if err != nil {
		return err
	}
	return nil
}

func write(r io.Reader, wr io.Writer) error {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return fmt.Errorf("ошибка чтения данных")
	}
	input := scanner.Text()
	words := strings.Fields(input)

	bufWriter := bufio.NewWriter(wr)

	for _, word := range words {
		if _, err := bufWriter.WriteString(word + "\n"); err != nil {
			return err
		}
	}
	return bufWriter.Flush()
}
