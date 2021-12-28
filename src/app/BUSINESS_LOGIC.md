# ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/events` folder for describe events
- `./app/models` folder for describe business models
- `./app/topics` folder for describe `const` topics