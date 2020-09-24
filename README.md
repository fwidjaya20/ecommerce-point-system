# E-Commerce Point System
User Reward Point System

##### Prerequisites
- Docker
- [Docker SH](https://github.com/fwidjaya20/docker-sh)
- PostgreSQL
- Golang
- Nats Streaming (Message Broker)

## Domain
- User Point

## Postman Documentation
- [API Documentation](https://www.getpostman.com/collections/49a719fb8171eb505d93)

## Solution
- Message Broker / Message Queue (Nats)

NATS merupakan salah satu Sistem Message Queue Open Source.

Untuk memisahkan context `Request` dan `Proses Logic`.

`Proses Request` hanya akan mengirim data point kedalam `Message Queue` (disimpan secara Immutable) dan akan mengembalikan `Request Id` untuk Inquiry.

`Proses Logic` memiliki worker untuk mengolah data yang ada di dalam `Message Queue`, sehingga kedua proses tersebut terpisah.

- Event Sourcing

Agar data Point bersifat Historical, sehingga mudah untuk melakukan proses `Tracking` ketika terdapat kesalahan pada data.

Memperepat proses agregat, sehingga hasil kalkulasi akan lebih cepat. 

## How to
### Run Nats Streaming
1. Create Bash File
```shell
vim nats_streaming

#!/usr/bin/env docker.sh

name=nats_streaming
image=nats-streaming:latest
opts="
	-p 4222:4222
	-p 8222:8222
"
args="
	-st SQL
	-msu 0
	-mm 0
	-mb 0
	--sql_driver postgres
	--sql_source 'postgres://postgres:password@172.17.0.2:5432/nats_persistence?sslmode=disable'
"
```
2. Grant Executable Access
```shell
chmod +x nats_streaming
```
3. Run Shell
```shell
./nats_streaming start
```
### Generate Nats Persistence Database
1. Create DB **nats_persistence**
2. Run this Query:
```SQL
CREATE TABLE IF NOT EXISTS ServerInfo (uniquerow INTEGER DEFAULT 1, id VARCHAR(1024), proto BYTEA, version INTEGER, PRIMARY KEY (uniquerow));
CREATE TABLE IF NOT EXISTS Clients (id VARCHAR(1024), hbinbox TEXT, PRIMARY KEY (id));
CREATE TABLE IF NOT EXISTS Channels (id INTEGER, name VARCHAR(1024) NOT NULL, maxseq BIGINT DEFAULT 0, maxmsgs INTEGER DEFAULT 0, maxbytes BIGINT DEFAULT 0, maxage BIGINT DEFAULT 0, deleted BOOL DEFAULT FALSE, PRIMARY KEY (id));
CREATE INDEX Idx_ChannelsName ON Channels (name(256));
CREATE TABLE IF NOT EXISTS Messages (id INTEGER, seq BIGINT, timestamp BIGINT, size INTEGER, data BYTEA, CONSTRAINT PK_MsgKey PRIMARY KEY(id, seq));
CREATE INDEX Idx_MsgsTimestamp ON Messages (timestamp);
CREATE TABLE IF NOT EXISTS Subscriptions (id INTEGER, subid BIGINT, lastsent BIGINT DEFAULT 0, proto BYTEA, deleted BOOL DEFAULT FALSE, CONSTRAINT PK_SubKey PRIMARY KEY(id, subid));
CREATE TABLE IF NOT EXISTS SubsPending (subid BIGINT, row BIGINT, seq BIGINT DEFAULT 0, lastsent BIGINT DEFAULT 0, pending BYTEA, acks BYTEA, CONSTRAINT PK_MsgPendingKey PRIMARY KEY(subid, row));
CREATE INDEX Idx_SubsPendingSeq ON SubsPending (seq);
CREATE TABLE IF NOT EXISTS StoreLock (id VARCHAR(30), tick BIGINT DEFAULT 0);

-- Updates for 0.10.0
ALTER TABLE Clients ADD proto BYTEA;
```
### Run App
```shell
go run ./cmd/app
```