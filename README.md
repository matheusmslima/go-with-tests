<h1 align="center">Learn Go with Tests — study repo</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/status-in%20progress-yellow?style=flat-square" />
  <a href="https://quii.gitbook.io/learn-go-with-tests"><img src="https://img.shields.io/badge/book-Learn%20Go%20with%20Tests-blue?style=flat-square" /></a>
</p>

## 📚 About

This repo holds my work through [**Learn Go with Tests**](https://quii.gitbook.io/learn-go-with-tests) by Chris James — a free, TDD-first introduction to Go where every concept is learned by writing a test first.

I'm going through it as part of a deliberate move from safety-critical aerospace software (flight control systems, DO-178C verification & validation) into **backend engineering with Go**, with a particular interest in **financial services / fintech**. The rigor around correctness that this book teaches through TDD is the same instinct I'm bringing over from that background.

Each chapter lives in its own folder, mirroring the book's structure. Every folder contains the working code and its test file, built up incrementally the way the book teaches it — red, green, refactor.

## ✅ Progress

### Go fundamentals

- [x] Install Go
- [x] Hello, world
- [x] Integers
- [x] Iteration
- [ ] Arrays and slices
- [ ] Structs, methods & interfaces
- [ ] Pointers & errors
- [ ] Maps
- [ ] Dependency Injection
- [ ] Mocking
- [ ] Concurrency
- [ ] Select
- [ ] Reflection
- [ ] Sync
- [ ] Context
- [ ] Intro to property based tests (Roman Numerals)
- [ ] Maths
- [ ] Reading files
- [ ] Templating
- [ ] Generics
- [ ] Revisiting arrays and slices with generics

### Build an application

- [ ] HTTP server
- [ ] JSON, routing and embedding
- [ ] IO and sorting
- [ ] Command line & project structure
- [ ] Time
- [ ] WebSockets

## 🛠️ Running the tests

Each chapter folder is self-contained. From inside a chapter's directory:

```bash
go test ./...
```

Or to run everything in the repo at once:

```bash
go test ./... -v
```

## 🎯 What's next

Once this is further along, the TDD habits and Go fundamentals from here feed directly into a separate project: a [concurrent order-matching engine](#) built to apply real Go concurrency patterns to a problem shape relevant to trading and payments systems.

## 🔗 Reference

- Book: [quii.gitbook.io/learn-go-with-tests](https://quii.gitbook.io/learn-go-with-tests)
- Source: [github.com/quii/learn-go-with-tests](https://github.com/quii/learn-go-with-tests)
