## bash

1. `-` behind a varabile
```shell
  if [[ -n ${KUBE_GIT_VERSION_FILE-} ]]; then
    kube::version::load_version_vars "${KUBE_GIT_VERSION_FILE}"
    return
  fi
```

> if the colon is omitted, the operator tests only for existence. (Bash Reference Manual - 3.5.3 Shell Parameter Expansion)
依旧不明白为什么要加`-`，经过测试，不加也是可行的。

2. `$Format:%%$'`
Magic???