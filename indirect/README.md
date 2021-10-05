# Indirect dependencies

Multiple times you need to load indirect dependencies, this are the ones that you
are not explicitly using but maybe one of your dependencies does.

Packages that are for tests within a dependency are not loaded, only those directly
used on your repository.
