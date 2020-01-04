// import org.jetbrains.kotlin.gradle.plugin.mpp.KotlinNativeTargetPreset

val publishingVersion = "0.0.1"
val publishingName = "CodegenExampleLib"

plugins {
    id("com.android.library") version Version.androidPlugin
    kotlin("multiplatform") version Version.kotlin
    id("kotlinx-serialization") version Version.kotlin
    id("com.moowork.node") version "1.3.1"
}

repositories {
    jcenter()
    mavenCentral()
    google()
}

android {
    compileSdkVersion(29)
    sourceSets {
        getByName("main") {
            manifest.srcFile("src/androidMain/AndroidManifest.xml")
        }
    }
}

node {
    version = "v12.11.1"
    npmVersion = "6.12.1"
}

kotlin {
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
            moduleKind = "umd"
            sourceMap = true
            nodejs()
        }
    }
    // TODO wasm
    // wasm32 {
    //    binaries.executable()
    // }

    // TODO bump ktor version for kotlin 1.3.61
    // iosTargets.forEach {
    //     targetFromPreset(presets[it] as KotlinNativeTargetPreset) {
    //         binaries.framework(publishingName, listOf(RELEASE, DEBUG)) {
    //             isStatic = true
    //         }
    //     }
    // }

    sourceSets {
        val commonMain by getting {
            dependencies {
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime-common:${Version.kotlinxSerialization}")
                implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core-common:${Version.kotlinxCoroutines}")
                implementation("io.ktor:ktor-client-core:${Version.ktor}")
            }
        }

        jvm().compilations["main"].defaultSourceSet {
            dependencies {
                implementation(kotlin("stdlib-jdk8"))
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime:${Version.kotlinxSerialization}")
                // TODO which ktor engine? https://ktor.io/clients/http-client/engines.html#jvm
                implementation("io.ktor:ktor-client-apache:${Version.ktor}")
            }
        }

        android {
            dependencies {
                implementation(kotlin("stdlib-jdk8"))
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime:${Version.kotlinxSerialization}")
                implementation("org.jetbrains.kotlinx:kotlinx-coroutines-android:${Version.kotlinxCoroutines}")
                implementation("io.ktor:ktor-client-android:${Version.ktor}")
            }
        }

        js().compilations["main"].defaultSourceSet {
            dependencies {
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime-js:${Version.kotlinxSerialization}")
                implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core-js:${Version.kotlinxCoroutines}")
                implementation("io.ktor:ktor-client-js:${Version.ktor}")
            }
        }

        val ios by creating {
            kotlin.srcDir("src/iosMain")
            dependencies {
                implementation("org.jetbrains.kotlinx:kotlinx-serialization-runtime-native:${Version.kotlinxSerialization}")
                implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core-native:${Version.kotlinxCoroutines}")
                implementation("io.ktor:ktor-client-ios:${Version.ktor}")
            }
        }

        // iosTargets.forEach {
        //     targetFromPreset(kotlin.presets[it] as KotlinNativeTargetPreset) {
        //         compilations["main"].defaultSourceSet {
        //             dependsOn(ios)
        //         }
        //     }
        // }
    }
}
