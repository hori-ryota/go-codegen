import org.gradle.api.internal.FeaturePreviews

enableFeaturePreview(FeaturePreviews.Feature.GRADLE_METADATA.name)

pluginManagement {
    repositories {
        google()
        gradlePluginPortal()
        jcenter()
    }
    resolutionStrategy {
        eachPlugin {
            if (requested.id.id.startsWith("com.android.")) {
                useModule("com.android.tools.build:gradle:${Version.androidPlugin}")
            }
            if (requested.id.id == "kotlinx-serialization") {
                useModule("org.jetbrains.kotlin:kotlin-serialization:${requested.version}")
            }
        }
    }
}
