<template>
  <div class="user-activity">
    <div class="post">
      <div class="user-block" />
      <blockquote>
        <p>hello，我叫周文喆，目前从事于物联网行业,不敢说是江南来的翩翩美少年，但绝不是惹人厌的油腻地中海。都说相逢即是缘，希望我们能有美好的际会，而不是世人般谩相识。</p>
        <h4>基本资料</h4><p>外文名：VingurZhou<br>中文名：周文喆<br>国籍：中国<br>出生地：江苏<br>生日：2000年3月<br>讨厌：蘑菇、肥肉、虫子、抗压吧的那群蛆虫和微博的那群脑残粉<br>喜欢：黑嘉嘉、Jennie Kim、周慧敏<br>爱好：MOBA、早茶、编程、撸自己猫<br>偶像：UZI(我的青春啊!)、刘星图、**（一位小众歌手）、星矶<br>代表作品：无<br>主要成就：无</p>
        <h4>其它资料</h4><p>毕业于不知名普本<br>养了只颜值于猫界天花板的布偶<br>偶然成为go-admin贡献者<br>成为appen组织成员<br>正在学习python和golang<br>会简单的爬虫、数据处理、可视化、分析<br>会简单的接口开发和敏捷开发后台（前端做不到定制化）</p>
      </blockquote>    </div>

    <div class="post">
      <div class="user-images">
        <el-carousel :interval="6000" type="card" height="220px">
          <el-carousel-item v-for="item in carouselImages" :key="item">
            <img :src="item+carouselPrefix" class="image">
          </el-carousel-item>
        </el-carousel>
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
        'https://raw.githubusercontent.com/Vingurzhou/hosting-sevice/main/blog/erweima/3641660314941_.pic.jpg',
        'https://raw.githubusercontent.com/Vingurzhou/hosting-sevice/main/blog/erweima/3651660314941_.pic.jpg',
        'https://raw.githubusercontent.com/Vingurzhou/hosting-sevice/main/blog/erweima/3661660314942_.pic.jpg'
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
