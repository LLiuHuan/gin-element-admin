<template>
  <div class="container" id="container" :class="{'right-panel-active': active}">
    <div class="form-container sign-up-container">
      <el-form :model="signUpData" :rules="signUpRules" ref="signUpRef" class="signUpForm">
        <h1>创建账户</h1>
        <div class="social-container">
          <a href="#" class="social">
            <svg class="icon" aria-hidden="true">
              <use xlink:href="#iconqq"></use>
            </svg>
          </a>
          <a href="#" class="social">
            <svg class="icon" aria-hidden="true">
              <use xlink:href="#iconweixin"></use>
            </svg>
          </a>
          <a href="#" class="social">
            <svg class="icon" aria-hidden="true">
              <use xlink:href="#icongithub"></use>
            </svg>
          </a>
        </div>
        <span>或使用您的账户</span>
        <el-form-item prop="name">
          <el-input prefix-icon="el-icon-s-custom" v-model="signUpData.name" placeholder="用户名"></el-input>
        </el-form-item>
        <el-form-item prop="pass">
          <el-input prefix-icon="el-icon-s-order" type="password" v-model="signUpData.pass" placeholder="密码"
                    autocomplete="off"
                    show-password></el-input>
        </el-form-item>
        <el-form-item prop="checkPass">
          <el-input prefix-icon="el-icon-s-order" type="password" v-model="signUpData.checkPass" placeholder="确认密码"
                    autocomplete="off"
                    show-password></el-input>
        </el-form-item>
        <el-form-item prop="captcha">
          <el-col :span="15">
            <el-input prefix-icon="el-icon-s-claim" v-model="signUpData.captcha.code" placeholder="验证码"></el-input>
          </el-col>
          <el-col :span="9">
            <img @click="refreshCaptcha" style="width: 100%; margin-left: 2px;" :src="signUpData.captcha.captchaImg"
                 alt="验证码"/>
          </el-col>
        </el-form-item>
        <button type="button" @click="signUpClick">注册</button>
      </el-form>
    </div>
    <div class="form-container sign-in-container">
      <el-form :model="signInData" :rules="signInRules" ref="signInRef" class="signInForm">
        <h1>登录</h1>
        <div class="social-container">
          <a href="#" class="social">
            <svg class="icon" aria-hidden="true">
              <use xlink:href="#iconqq"></use>
            </svg>
          </a>
          <a href="#" class="social">
            <svg class="icon" aria-hidden="true">
              <use xlink:href="#iconweixin"></use>
            </svg>
          </a>
          <a href="#" class="social">
            <svg class="icon" aria-hidden="true">
              <use xlink:href="#icongithub"></use>
            </svg>
          </a>
        </div>
        <span>或者使用你的帐户</span>
        <el-form-item prop="name">
          <el-input prefix-icon="el-icon-s-custom" v-model="signInData.name" placeholder="用户名"></el-input>
        </el-form-item>
        <el-form-item prop="pass">
          <el-input prefix-icon="el-icon-s-order" type="password" v-model="signInData.pass" placeholder="密码"
                    autocomplete="off"
                    show-password></el-input>
        </el-form-item>
        <el-form-item prop="captcha">
          <el-col :span="15">
            <el-input prefix-icon="el-icon-s-claim" v-model="signInData.captcha.code" placeholder="验证码"></el-input>
          </el-col>
          <el-col :span="9">
            <img @click="refreshCaptcha" style="width: 100%; margin-left: 2px;" :src="signInData.captcha.captchaImg"
                 alt="验证码"/>
          </el-col>
        </el-form-item>
        <a href="#">忘记密码了？</a>
        <button type="button" @click="signInClick">登录</button>
      </el-form>
    </div>
    <div class="overlay-container">
      <div class="overlay">
        <div class="overlay-panel overlay-left">
          <h1>Welcome Back!</h1>
          <p>To keep connected with us please login with your personal info</p>
          <button type="button" class="ghost" id="signIn" @click="setActive">登录</button>
        </div>
        <div class="overlay-panel overlay-right">
          <h1>Hello, Friend!</h1>
          <p>Enter your personal details and start journey with us</p>
          <button type="button" class="ghost" id="signUp" @click="setActive">注册</button>
        </div>
      </div>
    </div>
  </div>
  <footer>
    <p>
      Created with <i class="fa fa-heart"></i> by
      <a target="_blank" href="https://florin-pop.com">Florin Pop</a>
      - Read how I created this and how you can join the challenge
      <a target="_blank" href="https://www.florin-pop.com/blog/2019/03/double-slider-sign-in-up-form/">here</a>.
    </p>
  </footer>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, computed, onMounted, watch} from 'vue'
import {ElMessage} from "element-plus";

import {getCaptcha, login} from "../../apis/login"
import {useRoute, useRouter, LocationQuery} from 'vue-router'
import router from '../../router'
import {store} from "../../store";
import {UserMutationTypes} from "../../store/user/mutation-types";

export default defineComponent({
  name: '',
  setup() {
    // const router = useRouter()

    // ref
    const signUpRef = ref()
    const signInRef = ref()


    function getOtherQuery(query: LocationQuery) {
      return Object.keys(query).reduce((acc, cur) => {
        if (cur !== 'redirect') {
          acc[cur] = query[cur]
        }
        return acc
      }, {} as LocationQuery)
    }

    watch(useRoute(), current => {
      console.log(current.query)
      console.log(1111)
      if (current.query) {
        state.redirect = current.query.redirect?.toString() ?? ''
        state.otherQuery = getOtherQuery(current.query)
      }
    })

    // 数据校验
    const validatePass = (rule: any, value: any, callback: any) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== state.signUpData.pass) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };

    // 数据
    const state = reactive({
      active: ref<boolean>(false),
      signUpData: {
        name: "",
        pass: "",
        checkPass: "",
        captcha: {
          code: "",
          captchaImg: "",
          captchaId: ""
        },
      },
      signUpRules: {
        name: [
          {required: true, message: '请输入用户名', trigger: 'blur'},
        ],
        pass: [
          {required: true, message: '请输入密码', trigger: 'blur'},
        ],
        checkPass: [
          {required: true, message: '请输入确认密码', trigger: 'blur'},
          {validator: validatePass, trigger: 'blur'},
        ]
      },
      signInData: {
        name: "",
        pass: "",
        captcha: {
          code: "",
          captchaImg: "",
          captchaId: ""
        },
      },
      signInRules: {
        name: [
          {required: true, message: '请输入用户名', trigger: 'blur'},
        ],
        pass: [
          {required: true, message: '请输入密码', trigger: 'blur'},
        ]
      }
    })

    // region 注册方法
    const signUpClick = () => {
      ElMessage.info("暂不支持注册！")
    }
    // endregion

    // region 登录方法

    const signInClick = () => {
      // router.push('/dashboard1')
      // 整合数据
      let opt = {
        username: state.signInData.name,
        password: state.signInData.pass,
        captcha: state.signInData.captcha.code,
        captchaId: state.signInData.captcha.captchaId,
      }
      login(opt).then((res: any) => {
        if (res.code === 0) {
          ElMessage.success("登录成功！")
          refreshCaptcha()
          router.push({'path': '/redirect?dashboard1'}).catch(err => {
            console.warn(err)
          })
          store.commit(UserMutationTypes.SET_TOKEN, res.data.access_token)
        } else {
          ElMessage.error(res.msg)
        }
      })
    }

    // endregion

    // region 公共方法
    const setActive = () => {
      state.active = !state.active
    }

    const refreshCaptcha = () => {
      getCaptcha().then((res: any) => {
        if (res.code === 0) {
          if (state.active) {
            state.signUpData.captcha.captchaImg = res.data.picPath;
            state.signUpData.captcha.captchaId = res.data.captchaId;
          } else {
            state.signInData.captcha.captchaImg = res.data.picPath;
            state.signInData.captcha.captchaId = res.data.captchaId;
          }
        } else {
          ElMessage.error(res.msg)
        }
      })
    }

    onMounted(() => {
      refreshCaptcha()
    })

    // endregion
    return {
      ...toRefs(state),
      setActive,

      signUpRef,
      signUpClick,

      signInRef,
      signInClick,

      refreshCaptcha,
    }
  }
})
</script>

<style lang="scss" scoped>
//@import url("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.14.0/css/all.min.css");
//@import url('https://fonts.googleapis.com/css?family=Montserrat:400,800');

@import url('src/styles/login.scss');
</style>