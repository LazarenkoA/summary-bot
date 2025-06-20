#!/bin/bash

# Проверка наличия переменной окружения
if [ -z "$REMOTE_IP" ]; then
  echo "Ошибка: переменная окружения REMOTE_IP не установлена."
  exit 1
fi

# Сборка бинарника
echo "Сборка Go-программы..."
CGO_ENABLED=0 go build -o summary ./main.go

# Загрузка по SFTP
sftp -i ~/.ssh/key artem@"$REMOTE_IP" <<EOF
put summary /var/tmp/
exit
EOF

echo "Готово. md5sum - $(md5sum summary)"
