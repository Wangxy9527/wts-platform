<template>
<div>
  <el-card class="box-card">
    <div>
      <span>我的足迹</span>

    </div>
  </el-card>

</div>
</template>

<script>
export default {
  name: 'InfoFoot',
  data () {
    return {
      ruleForm: {
        username: '',
        declaration: ''
      }
    }
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          alert('submit!')
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style scoped>
  .box-card {
    width: 97%;
    margin: auto;
    float: right;
  }
</style>
