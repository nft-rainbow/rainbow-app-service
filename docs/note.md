# Dodo

## 获得群主身份信息

[获取群信息](https://open.imdodo.com/dev/api/island.html#%E8%8E%B7%E5%8F%96%E7%BE%A4%E4%BF%A1%E6%81%AF)-> 响应中取ownerDodoSourceId -> [获取成员信息](https://open.imdodo.com/dev/api/member.html#%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98%E4%BF%A1%E6%81%AF)

## 通过dodo号获取dodo source id
[获取成员dodo号映射列表](https://open.imdodo.com/dev/api/member.html#%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98dodo%E5%8F%B7%E6%98%A0%E5%B0%84%E5%88%97%E8%A1%A8)


# 数据结构

- SocialToolServer: {[]PushInfo, RainbowUser}
- PushInfo: {Activity, Channel, IdentityGroup, Msg, ColorTheme}
