<template>
  <div class="user-activity">
    <div class="post">
      <div class="user-block">
        <img class="img-circle" :src="'https://s1.aigei.com/src/img/png/0d/0d531e281cb24816a88a1e1a75e5cfa7.png?imageMogr2/auto-orient/thumbnail/!234x234r/gravity/Center/crop/234x234/quality/85/&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:KHMz3zP4c378yDWtLjW6bBoDRBQ='">
        <span class="username text-muted"><a href="tel:19952429930">Telephone</a></span>
        <span class="description"><a href="tel:19952429930">19952429930</a></span>
      </div>

    </div>
    <div class="post">
      <div class="user-block">
        <img class="img-circle" :src="'https://s1.aigei.com/src/img/png/04/04bb64ab090d4e83930fb724c7896386.png?imageMogr2/auto-orient/thumbnail/!234x234r/gravity/Center/crop/234x234/quality/85/&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:nQh9Hz2w-bmWJ8ph1mf3qEJWhaU='">
        <span class="username text-muted"><a href="mailto:zwz0123460218@icloud.com">Email</a></span>
        <span class="description"><a href="mailto:zwz0123460218@icloud.com">zwz0123460218@icloud.com</a></span>
      </div>

    </div>
    <div class="post">
      <div class="user-block">
        <img class="img-circle" :src="'https://s1.aigei.com/src/img/png/57/5767c3d9931d40aabbf1d15e769585b1.png?imageMogr2/auto-orient/thumbnail/!234x234r/gravity/Center/crop/234x234/quality/85/&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:vlix16T5fHUMS8fALOuQTa1-dUE='">
        <span class="username text-muted"><a href="weixin://">WeChat</a></span>
        <span class="description"><a href="weixin://">zwz786000694</a></span>
      </div>

    </div>
    <div class="post">
      <div class="user-block">
        <img class="img-circle" :src="'https://s1.aigei.com/src/img/png/0d/0de8a065a2cf46689e3f24245deaa5df.png?imageMogr2/auto-orient/thumbnail/!234x234r/gravity/Center/crop/234x234/quality/85/&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:xgemesfzTSGimcIrzzim7AZWETs='">
        <span class="username text-muted">QQ</span>
        <span class="description">786000694</span>
      </div>
    </div>
    <div class="post">
      <div class="user-block">
        <img class="img-circle" :src="'https://s1.aigei.com/src/img/png/a3/a3d87bbda5554e5ca26463f03ed6f5be.png?imageMogr2/auto-orient/thumbnail/!234x234r/gravity/Center/crop/234x234/quality/85/&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:-uVBcfyMwHt7Q-C4hWb3sZPQwZY='">
        <span class="username text-muted">Facebook</span>
        <span class="description">vingurzhou</span>
      </div>
    </div>
    <div class="post">
      <div class="user-block">
        <img class="img-circle" :src="'https://s1.aigei.com/src/img/png/c0/c005097d3b22424fb9909fd2cca24575.png?imageMogr2/auto-orient/thumbnail/!234x234r/gravity/Center/crop/234x234/quality/85/&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:pBVOKDX4Z7X2x1wBuuwNdJaySu0='">
        <span class="username text-muted">WeiBo</span>
        <span class="description">vingurzhou</span>
      </div>
    </div>

    <div class="post">
      <div class="user-block">
        <img class="img-circle" :src="'https://s1.aigei.com/src/img/png/81/81d9d708b81948f2b72cb682a0851dcc.png?imageMogr2/auto-orient/thumbnail/!234x234r/gravity/Center/crop/234x234/quality/85/&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:84qbqhVeL4ftE8gW29ekcgmiB6g='">
        <span class="username text-muted">Twitter</span>
        <span class="description">vingurzhou</span>
      </div>

    </div>
    <div class="post">
      <div class="user-block">
        <img class="img-circle" :src="'https://s1.aigei.com/src/img/png/b3/b3a6079d1bc148bcb8e561434c257b66.png?imageMogr2/auto-orient/thumbnail/!234x234r/gravity/Center/crop/234x234/quality/85/&e=1735488000&token=P7S2Xpzfz11vAkASLTkfHN7Fw-oOZBecqeJaxypL:NdhVuWmN2VO-YJPS4Rk44dlAt3o='">
        <span class="username text-muted">Instagram</span>
        <span class="description">vingurzhou</span>
      </div>

    </div>
  </div>
</template>

<script>
const avatarPrefix = '?imageView2/1/w/80/h/80'
const carouselPrefix = '?imageView2/2/h/440'

export default {
  data() {
    return {
      carouselImages: [
        'https://wpimg.wallstcn.com/9679ffb0-9e0b-4451-9916-e21992218054.jpg',
        'https://wpimg.wallstcn.com/bcce3734-0837-4b9f-9261-351ef384f75a.jpg',
        'https://wpimg.wallstcn.com/d1d7b033-d75e-4cd6-ae39-fcd5f1c0a7c5.jpg',
        'https://wpimg.wallstcn.com/50530061-851b-4ca5-9dc5-2fead928a939.jpg'
      ],
      avatarPrefix,
      carouselPrefix
    }
  }
}
</script>

<style lang="scss" scoped>
.user-activity {
  .user-block {

    .username,
    .description {
      display: block;
      margin-left: 50px;
      padding: 2px 0;
    }

    .username{
      font-size: 16px;
      color: #000;
    }

    :after {
      clear: both;
    }

    .img-circle {
      border-radius: 50%;
      width: 40px;
      height: 40px;
      float: left;
    }

    span {
      font-weight: 500;
      font-size: 12px;
    }
  }

  .post {
    font-size: 14px;
    border-bottom: 1px solid #d2d6de;
    margin-bottom: 15px;
    padding-bottom: 15px;
    color: #666;

    .image {
      width: 100%;
      height: 100%;

    }

    .user-images {
      padding-top: 20px;
    }
  }

  .list-inline {
    padding-left: 0;
    margin-left: -5px;
    list-style: none;

    li {
      display: inline-block;
      padding-right: 5px;
      padding-left: 5px;
      font-size: 13px;
    }

    .link-black {

      &:hover,
      &:focus {
        color: #999;
      }
    }
  }

}

.box-center {
  margin: 0 auto;
  display: table;
}

.text-muted {
  color: #777;
}
</style>
