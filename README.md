# TeaBuf 3

「茶葉緩衝 3」是基於「緩衝協議 3」（Protocol Buffer 3）的功能加強版本，最終會完美地轉譯成原生的 `.proto` 檔案。

## VS Code 程式碼螢光標記

茶葉緩衝 3 支援 VS Code 的程式碼螢光標記（Syntax Highlight），可以至 [Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=YamiOdymel.vscode-teabuf3) 下載。

## 功能改進

### 省去分號

```proto
syntax = "proto3";
package pb;

message User {
    // ...
}
```

```proto
package pb

message User {
    // ...
}
```

### 引用匯入

```proto
syntax = "proto3";
package pb;
import "another.proto";
```

```proto
package pb
import ../../my/another.proto
```

### 省去編號欄位、型態倒置

```proto
message User {
    uint64 id       = 1;
    string username = 2;
}
```

```proto
message User {
    id       uint64
    username string
}
```

### 型態簡略

```proto
message User {
    uint64 id    = 1;
    uint64 count = 2;
}
```

```proto
message User {
    id    uint
    count uint
}
```

### 型態別名

```proto
message User {
    string icon_url   = 1;
    string cover_url  = 2;
    string avatar_url = 3;
}
```

```proto
type URL string

message User {
    icon_url   URL
    cover_url  URL
    avatar_url URL
}
```

### 訊息擴展

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

```proto
message Time {
    created_at uint
    deleted_at uint
    updated_at uint
}

message User {
    Time

    id       uint
    username string
}
```

### 重複欄位

```proto
message Category {
    repeated uint64 ids = 1;
}
```

```proto
message Category {
    ids []uint
}
```

### 服務定義

```proto
service UserService {
    rpc CreateUser (CreateUserRequest) returns (CreateUserResponse) {}
    rpc DeleteUser (DeleteUserRequest) returns (DeleteUserResponse) {}
}
```

```proto
service UserService {
    CreateUser(CreateUserRequest) CreateUserResponse {}
    DeleteUser(DeleteUserRequest) DeleteUserResponse {}
}
```