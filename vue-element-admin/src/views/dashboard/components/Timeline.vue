<template>
  <div class="block">
    <el-timeline>
      <el-timeline-item v-for="(item,index) of timeline" :key="index" :timestamp="item.timestamp" placement="top">
        <el-card>
          <h4>{{ item.title }}</h4>
          <p>{{ item.content }}</p>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </div>
</template>

<script>
export default {
  data() {
    return {
      timeline: [
        {
          timestamp: '2022/8/10',
          title: '关于我成为go-admin贡献者这件事',
          content: 'Vingurzhou committed 2019/4/21 20:46'
        },
        {
          timestamp: '2022/8/12',
          title: '关于我养了只超靓的布偶这件事',
          content: 'Vingurzhou committed 2019/4/20 20:46'
        },
        {
          timestamp: '2021/5/22',
          title: '关于我写的小说被盗版这件事',
          content: 'Vingurzhou committed 2019/4/22 20:46'
        },
        {
          timestamp: '2021/3/23',
          title: '关于我在比心做陪玩这件事',
          content: 'Vingurzhou committed 2019/4/23 20:46'
        }
      ]
    }
  }
}
</script>
