@echo off
echo Building PrismaController.
packr2 clean
packr2
go build -ldflags -H=windowsgui .
packr2 clean
cls
echo Prisma Controller Builded.