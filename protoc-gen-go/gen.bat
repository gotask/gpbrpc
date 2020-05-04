@echo off
for %%c in (./*.proto) do (
	echo %%c
	protoc.exe --go_out=plugins=gpbrpc:. %%c
)
pause