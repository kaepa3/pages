---
title: "Build_GPU_environment"
date: 2019-08-24T08:49:46+09:00
draft: true
---

{{< talk  >}}
会社の人間が取りあえずAIとか言ってくるから出家して欲しい
{{< /talk >}}

deep learning始めました
というわけで、ゼロから作るDeep Learning2をやってるんですが
会社のPCはGPU付きのデスクトップが有るんですね。
じゃあ、動かして見ようじゃないって話ですよ。

先ずはエラーが発生します。今回はこいつをどうにかしようが主な内容です。
```
ModuleNotFoundError: No Module 'cupy'
```

cupyなんか入れてないPCなので先ずはcudaのバージョンを確認するためにtoolkitをインストールします。
(ここから)[https://developer.nvidia.com/cuda-toolkit]

ver10.1が確認できました。

対応するcupyを追加しますがエラーです。

オフィシャルの通りpipの更新を行います。結果ダメです。

pip install cupy でエラーでも確認するかと試みます。


止まったね、情報は渡さないとのことかと思い止めたら
VS2017のC++ビルドツールをいれろとの事で入れる。
なお、VS2017は入ってる環境なのでvisual studio installer から追加。

なおも動かない、vagrantからLinuxにしようかと思って退陣を始める。
その中でpipのget_supported()を実行した結果win32で動いている
、、、、、、、

64bitのpythonをインストール。

ハイ成功、すいませんwinの悪口を頭の中で言ってました。

で、実行。エラー、、、、


atなんかないそうな、こっちは写経しているのにと思ったら
マジで参考書の方のミス（バージョン？）みたいなのでコメントアウトしてソース直す
[参考](https://teratail.com/questions/197824)

動作を確認。。。。え、GPU使ってもそこそこ時間いるやん。

