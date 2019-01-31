# TeaBuf 3

「茶葉緩衝 3」是基於「緩衝協議 3」（Protocol Buffer 3）的功能加強版本。最終會由 Golang 所撰寫的轉譯器，完美地將茶葉緩衝轉譯成原生的 `.proto` 檔案。

## VS Code 程式碼螢光標記

茶葉緩衝 3 支援 VS Code 的程式碼螢光標記（Syntax Highlight），可以至 [Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=YamiOdymel.vscode-teabuf3) 下載。

## 功能改進

* [省去分號](#省去分號)
* [引用匯入](#引用匯入)
* [省略欄位編號](#省略欄位編號)
* [型態倒置](#型態倒置)
* [型態簡略](#型態簡略)
* [型態別名](#型態別名)
* [訊息擴展](#訊息擴展)
* [重複欄位](#重複欄位)
* [服務定義](#服務定義)

### 省去分號

**Protocol Buffer**：每個語句的最後都需要加上分號。

```proto
syntax = "proto3";
package pb;

message User {
    // ...
}
```

**Tea Buffer**：語句最後的分號可以省略，且開頭不需要 `syntax = "proto3";` 聲明。

```proto
package pb

message User {
    // ...
}
```

### 引用匯入

**Protocol Buffer**：檔案引用不可使用相對路徑，且須加引號包圍。

```proto
syntax = "proto3";
package pb;
import "another.proto";
```

**Tea Buffer**：檔案可以透過相對路徑相互引用並且無須雙引號。

```proto
package pb
import ../../my/another.proto
```

### 省略欄位編號

**Protocol Buffer**：欄位皆需透過編號進行二進制排序定義。

```proto
message User {
    uint64 id       = 1;
    string username = 2;
}
```

**Tea Buffer**：欄位會以程式自動依照順序排定，無須手動指定。

```proto
message User {
    id       uint64
    username string
}
```

### 型態倒置

**Protocol Buffer**：型態在前，欄位名稱在後。

```proto
message User {
    uint64 id       = 1;
    string username = 2;
}
```

**Tea Buffer**：型態在後，更能第一時間看出所有可用欄位，且更符合 Golang 原生用法與習慣。

```proto
message User {
    id       uint64
    username string
}
```

### 型態簡略

**Protocol Buffer**：需要特別指定 64 或是 32 位元的正整數。

```proto
message User {
    uint64 id    = 1;
    uint64 count = 2;
}
```

**Tea Buffer**：任何正整數預設都是 64 位元。

```proto
message User {
    id    uint
    count uint
}
```

### 型態別名

**Protocol Buffer**：型態無法有別名，僅能透過無意義的 `string` 或 `uint64` 作為代表（請參閱：[Provide support for type aliases](https://github.com/protocolbuffers/protobuf/issues/3521)）。

```proto
message User {
    string icon_url   = 1;
    string cover_url  = 2;
    string avatar_url = 3;
}
```

**Tea Buffer**：可以自訂型態別名來讓文件中的每個欄位都更有意義。

```proto
type URL string

message User {
    icon_url   URL
    cover_url  URL
    avatar_url URL
}
```

### 訊息擴展

**Protocol Buffer**：需要複用某些欄位時，必須手動複製無法延展或是繼承現有訊息定義（請參閱：[Extending Protobuf Messages](https://stackoverflow.com/questions/29263507/extending-protobuf-messages)）。

```proto
message Time {
    uint64 created_at = 1;
    uint64 deleted_at = 2;
    uint64 updated_at = 3;
}

message User {
    uint64 id         = 1;
    string username   = 2;
    uint64 created_at = 3;
    uint64 deleted_at = 4;
    uint64 updated_at = 5;
}
```

**Tea Buffer**：直接將現有訊息定義擺放於另一個訊息中即可擴展、暴露該欄位定義而無須手動複製。

```proto
message Time {
    created_at uint
    deleted_at uint
    updated_at uint
}

message User {
    id       uint
    username string

    Time
}
```

### 重複欄位

**Protocol Buffer**：透過額外的 `repeated` 關鍵字表明某個欄位是陣列、會重複。

```proto
message Category {
    repeated uint64 ids = 1;
}
```

**Tea Buffer**：直接以較直覺、類程式的方式定義一個會重複的陣列欄位型態。

```proto
message Category {
    ids []uint
}
```

### 服務定義

**Protocol Buffer**：多餘的服務方法定義關鍵字（如：`rpc`、`returns`）。

```proto
service UserService {
    rpc CreateUser (CreateUserRequest) returns (CreateUserResponse) {}
    rpc DeleteUser (DeleteUserRequest) returns (DeleteUserResponse) {}
}
```

**Tea Buffer**：省去多餘關鍵字，類似定義程式的 Interface 介面，並且最終的區塊包覆符號可以省略。

```proto
service UserService {
    CreateUser(CreateUserRequest) CreateUserResponse
    DeleteUser(DeleteUserRequest) DeleteUserResponse
}
```