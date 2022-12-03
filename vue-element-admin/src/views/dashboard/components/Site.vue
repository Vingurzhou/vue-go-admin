<template>
  <div class="user-activity">
    <div class="post">
      <div class="user-block" />
      <blockquote>
        <h4>关于本站</h4>
        <p> .....说来话长，这不是我第一个博客网站，大学期间曾经基于hexo主题做了个唯美如画的博客网站（是真的很好看= =），但毕竟学的后端语言，之前的博客虽然唯美好看，却是静态页面，而且记笔记还得markdown写，写完再用hexo编译渲染，属实麻烦。</p>
        <p>没有办法，我就新建了这个博客网站，前端是基于vue-element-admin的，后端是基于gin的。虽然前后端分离开发貌似也能在服务器的两个端口上，但是我还是想将服务放在不同的地方，但是我又穷...于是我将前端置于github page上，后端置于自己的一台阿里云服务器上，至于存储服务除了数据库和个别,其它的我放到了一个github repository里。</p>
        <p>肯定有人要说，哎呀呀，博客网站也算项目啊。一个博客确实不能算啥项目，但我还未开发完，之后我会在此站台加上网关的开发、负载均衡、还有各种开发的工具都嵌入进来。</p>
        <p>好奇者肯定已经在此站点击各个模块了，但是会发现有些页面加载不出来，主要是因为后端接口我开发都只是在本地测试，还没在服务器打开🙁。</p>
        <p>总之这个小站肯定没有之前的好看的，但主要注重功能嘛，就是记记自己的所爱所想，开发些方便自己的工具，希望各位访客不要介意。</p>
      </blockquote>    </div>

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
