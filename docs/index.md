# Welcome to Reward's documentation!

``` include:: ../README.md
    :start-after: <!-- include_open_start -->
    :end-before: <!-- include_open_stop -->
```

## Features

``` include:: ../README.md
    :start-after: Features
    :end-before: <!-- include_open_stop -->

```

``` toctree::
   :maxdepth: 2
   :caption: Contents:
   :glob:

   installation
   getting-started
   services
   environments
   usage
   configuration
   autocompletion
   faq
   experimental
```

``` toctree::
   :maxdepth: 2
   :caption: About:
   :glob:

   changelog
```

Under the hood `docker-compose` is used to control everything which Reward runs
(shared services as well as per-project containers) via the Docker Engine.


## Acknowledgement

``` include:: ../README.md
    :start-after: Acknowledgement
    :end-before: <!-- include_open_stop -->

```

* [David Alger's Github](https://github.com/davidalger)
* [Warden's homepage](https://warden.dev)
