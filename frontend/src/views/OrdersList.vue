<template>
  <el-row>

    <el-col :md="4" :sm="8" :lg="3">
      <div class="customHeader">
        <div class="searchLabel">
          <font-awesome-icon icon="fa-solid fa-magnifying-glass" class="fa-search-icon" size="3x"/>
          <label>Search</label>
        </div>
      </div>
    </el-col>
    <el-col :sm="16" :md="20">
      <el-input>
        
      </el-input>
    </el-col>
  </el-row>
</template>

<script lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

export default {
  components: { FontAwesomeIcon },

}
</script>

<style lang="scss" scoped>

.searchLabel {
  font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
  font-weight: 600;
  text-align: center;
}

.searchLabel {
  label {
    font-size: 2em;
    font-weight: bold;
  }

  .fa-search-icon {
    width: 30px;
    height: 30px;
    line-height: 12px;
    margin-left: 8px;
    margin-right: 8px;
  }
}
</style>
