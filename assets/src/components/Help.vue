<template>
  <div class="p-4">
    <div class="pb-4">
      <h2 class="h6 text-black-50 font-weight-bold">■ Knowmeとは</h2>
      <p>
        社内のコミュニケーション活性化を目指して、社員同士の食事会の料金の一部を補助する社内制度です。
      </p>
    </div>
    
    <div class="pb-4">
      <h2 class="h6 text-black-50 font-weight-bold">■ 対象者</h2>
      <p>@candee.co.jpを持つ全社員</p>
    </div>

    <div class="pb-4">
      <h2 class="h6 text-black-50 font-weight-bold">■ ルール</h2>
      <ul>
        <li>@candee.co.jpでサインインしてください</li>
        <li>3人1組(2部署以上必須、4人以上NG)で行ってください</li>
        <li>1人あたり月1回まで利用可能です</li>
        <li>1度でも一緒に行った方とは再度行けません</li>
        <li>同日に昼夜1組づつ利用可能です</li>
        <li>
          <strong style="text-decoration: underline;">
            開催後、翌営業日に3人が写っている楽しい写真をアップロードすること♪
          </strong>
        </li>
      </ul>
    </div>

    <div class="pb-4">
      <h2 class="h6 text-black-50 font-weight-bold">■ 補助内容</h2>
      <dl>
        <dt><badge-day/></dt><dd>1620円(税込)/名[毎月先着4組]</dd>
        <dt><badge-night/></dt><dd>3240円(税込)/名[毎月先着9組]</dd>
      </dl>
    </div>

    <div class="pb-4">
      <h2 class="h6 text-black-50 font-weight-bold">■ 決済方法</h2>
      <p>
        立替申請なCandee社員代表者が全員部まとめて立替申請をしてください。費用は給与と一緒に振り込まれます。
      </p>
      <h3 class="h6 mt-2">【領収書の内容】</h3>
      <div class="border p-3" style="width: 18rem;">
        <dl>
          <dt>宛名</dt><dd>(株)Candee</dd>
          <dt>料金[昼の場合]</dt><dd>4,860円(税込み)</dd>
          <dt>料金[夜の場合]</dt><dd>9,720円(税込み)</dd>
        </dl>
      </div>
    </div>

    <div class="pb-4">
      <h2 class="h6 text-black-50 font-weight-bold">■ 注意点</h2>
      <ul>
        <li>実施日の前日までに等サイトで申請してください</li>
        <li>実施日には必ず申請したメンバーで開催してください</li>
        <li>事後申請はできません</li>
        <li>来月以降の登録はできません</li>
      </ul>
    </div>
  </div>
</template>

<script>
import BadgeDay from './BadgeDay.vue'
import BadgeNight from './BadgeNight.vue'

export default {
  name: 'Help',
  components: {
    'badge-day': BadgeDay,
    'badge-night': BadgeNight,
  },
}
</script>
