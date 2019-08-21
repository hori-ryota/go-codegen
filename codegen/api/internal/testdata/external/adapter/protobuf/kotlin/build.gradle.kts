import org.jetbrains.kotlin.gradle.plugin.mpp.KotlinNativeTargetPreset

plugins {
    id("com.android.library")
    kotlin("multiplatform") version Version.kotlin
    id("kotlinx-serialization") version Version.kotlin
}

repositories {
    jcenter()
    mavenCentral()
    google()
}

android {
    compileSdkVersion(28)
    sourceSets {
        getByName("main") {
            manifest.srcFile("src/androidMain/AndroidManifest.xml")
        }
    }
}

kotlin {
    val libName = "CodegenExampleLib"

    val iosTargets = listOf("iosArm32", "iosArm64", "iosX64")

    jvm {
        compilations["main"].kotlinOptions {
            jvmTarget = "1.8"
        }
    }
    android()

    js {
        compilations["main"].kotlinOptions {
            main = "noCall"
            target = "v5"
        }
    }
    // TODO wasm
    // wasm32 {
    //    binaries.executable()
    // }

    iosTargets.forEach {
        targetFromPreset(presets[it] as KotlinNativeTargetPreset) {
            binaries.framework(libName, listOf(RELEASE, DEBUG)) {
                isStatic = true
            }
        }
    }

    sourceSets {
        val commonMain by getting {
            dependencies {
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime-common:${Version.kotlinxSerialization}")
            }
        }

        jvm().compilations["main"].defaultSourceSet {
            dependencies {
                implementation(kotlin("stdlib-jdk8"))
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime:${Version.kotlinxSerialization}")
            }
        }

        android {
            dependencies {
                implementation(kotlin("stdlib-jdk8"))
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime:${Version.kotlinxSerialization}")
            }
        }

        js().compilations["main"].defaultSourceSet {
            dependencies {
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime-js:${Version.kotlinxSerialization}")
            }
        }

        val ios by creating {
            kotlin.srcDir("src/iosMain")
            dependencies {
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime-native:${Version.kotlinxSerialization}")
            }
        }

        iosTargets.forEach {
            targetFromPreset(kotlin.presets[it] as KotlinNativeTargetPreset) {
                compilations["main"].defaultSourceSet {
                    dependsOn(ios)
                }
            }
        }
    }
}
