package config

import (
	"bytes"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputDomains(t *testing.T) {
	tmpPath := path.Join(os.TempDir(), "test_conf.txt")
	defer os.Remove(tmpPath)
	text := []byte("test.test test1.test")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(text)
	assert.NoError(t, err, "ошибка при добавлении текста в инпут")
	w.Close()

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	os.Stdin = r

	err = InputDomains(tmpPath)
	assert.NoError(t, err, "ошибка записи")

	content, err := os.ReadFile(tmpPath)
	assert.NoError(t, err, "ошибка при чтении файла")

	expected := `test.test
test1.test
`

	assert.Equal(t, expected, string(content), "содержание файла не совпадает")
}

func TestWrite(t *testing.T) {
	text := "test.test test1.test"
	reader := strings.NewReader(text)
	var buffer bytes.Buffer
	err := write(reader, &buffer)
	assert.NoError(t, err, "ошибка чтения")
	res := buffer.String()
	reader = strings.NewReader(res)
	domains, err := parse(reader)
	assert.NoError(t, err, "ошибка парсинга")
	assert.Equal(t, 2, len(domains), "ошибка количества полученных доменов")
}

func TestParse(t *testing.T) {
	testDomains := `test.test
	test1.test
	test2.test`
	reader := strings.NewReader(testDomains)

	domains, err := parse(reader)
	assert.NoError(t, err, "ошибка парсинга")
	assert.Equal(t, 3, len(domains), "ошибка количества полученных доменов")
}

func TestNewConfig(t *testing.T) {
	tmpPath := path.Join(os.TempDir(), "test_conf.txt")
	defer os.Remove(tmpPath)
	domains, err := Domains(tmpPath)
	assert.Error(t, err, "нет ошибки при отсутствующем файле конфигурации")
	testData := `test.test
	test1.test
	`
	err = os.WriteFile(tmpPath, []byte(testData), 0644)
	assert.NoError(t, err, "ошибка при создании временного файла")

	domains, err = Domains(tmpPath)
	assert.NoError(t, err, "ошибка при парсинге временного файла")
	assert.Equal(t, 2, len(domains), "ошибка количества полученных доменов")
	if len(domains) >= 1 {
		assert.Equal(t, "test.test", domains[0].Url, "неверные данные в конфиге")
		if len(domains) > 1 {
			assert.Equal(t, "test1.test", domains[1].Url, "неверные данные в конфиге")
		}
	}

}
