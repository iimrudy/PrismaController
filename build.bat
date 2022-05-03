@echo off
echo Building PrismaController.
go build -ldflags -H=windowsgui .
cls
echo Prisma Controller Compiled.