# GO 言語の基礎

## ファイル分割について

Go 言語では 1 つのディレクトリに 1 つのパッケージのみを作成できる。  
1 つのパッケージ(ディレクトリ)内において、ディレクトリ内のソースファイルは複数のファイルに分割することができる。  
分割した複数のファイル同士は定義された変数や関数などを共有する。

以下の階層ツリーを例に説明を行う。

```
.
├── go.mod
├── main.go
├── test
│   ├── sample.go
│   └── sample_test.go
├── xxx
│   ├── xxx.go
│   └── zzz.go
└── yyy
    └── yyy.go
```

### メインモジュールについて

対象

- `main.go`
- `go.mod`

アプリケーションのメイン関数となるソースファイル、およびモジュール定義ファイルはプロジェクトルートに置いておく。  
この状態で

```bash
go run main.go
```

を実行すればアプリケーションを実行することができる。

### 自作パッケージについて

対象

- `xxx/*`
- `yyy/*`
- `test/*`

これらはそれぞれ自作パッケージとなっており、`main.go`から import することができる。  
パッケージ内で定義された変数や関数は Private 属性や Public 属性を持つ。

- Public 属性

パッケージ外部から直接参照することができる。  
定義した変数名や関数名が大文字から始まっていると Public 属性となる。

- Private 属性

パッケージ外部から直接参照することはできないが、同一パッケージ内の異なるファイルから直接参照できる。  
定義した変数名や関数名が小文字から始まっていると Private 属性となる。

```go
/* xxx/xxx.go */

// Public属性
// main.goでは、xxxをimportすれば直接参照できる。
// xxx/zzz.goでは、何もしなくても直接参照できる。
const PublicVariable = "Public"

// Private属性
// main.goでは、xxxをimportしても直接参照できない。
// xxx/zzz.goでは、何もしなくても直接参照できる。
const privateVariable = "Private"
```

### テストについて

対象

- `test/sample.go`
- `test/sample_test.go`

Go 言語のテストコードファイルは、対象のファイル名に`_test`のサフィックスつける。

`sample.go` -> `sample_test.go`

# 参考

https://pod.hatenablog.com/entry/2018/12/26/074944
