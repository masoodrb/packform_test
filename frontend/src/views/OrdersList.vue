<template>
  <el-row>

    <el-col :span="4">
      <div class="customHeader">
        <faUserSecret></faUserSecret>
        <div class="searchLabel">
            <Search width="40" height="40" />
            <label>Search</label>
        </div>
      </div>
    </el-col>
    <el-col :span="20">
      <el-input>
        
      </el-input>
    </el-col>
  </el-row>
</template>

<script lang="ts">
// import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faUserSecret } from '@fortawesome/free-solid-svg-icons'
export default {
  components: { faUserSecret },

}
</script>

<style lang="scss" scoped>

.searchLabel {
  font-family: 'Helvetica Neue';
  font-weight: 600;
}

.searchLabel {
  label {
    font-size: 2em;
    font-weight: bold;
  }
}
</style>
