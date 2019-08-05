// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: codegen.proto

package com.github.hori-ryota/go-codegen/api/example/codegen;

/**
 * Protobuf type {@code codegen.DoingSomethingWithoutOutputAndActorUsecaseInput}
 */
public  final class DoingSomethingWithoutOutputAndActorUsecaseInput extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:codegen.DoingSomethingWithoutOutputAndActorUsecaseInput)
    DoingSomethingWithoutOutputAndActorUsecaseInputOrBuilder {
private static final long serialVersionUID = 0L;
  // Use DoingSomethingWithoutOutputAndActorUsecaseInput.newBuilder() to construct.
  private DoingSomethingWithoutOutputAndActorUsecaseInput(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private DoingSomethingWithoutOutputAndActorUsecaseInput() {
    stringParam_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new DoingSomethingWithoutOutputAndActorUsecaseInput();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private DoingSomethingWithoutOutputAndActorUsecaseInput(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();

            stringParam_ = s;
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.github.hori-ryota/go-codegen/api/example/codegen.CodegenProto.internal_static_codegen_DoingSomethingWithoutOutputAndActorUsecaseInput_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.github.hori-ryota/go-codegen/api/example/codegen.CodegenProto.internal_static_codegen_DoingSomethingWithoutOutputAndActorUsecaseInput_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput.class, com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput.Builder.class);
  }

  public static final int STRING_PARAM_FIELD_NUMBER = 1;
  private volatile java.lang.Object stringParam_;
  /**
   * <code>string string_param = 1;</code>
   */
  public java.lang.String getStringParam() {
    java.lang.Object ref = stringParam_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      stringParam_ = s;
      return s;
    }
  }
  /**
   * <code>string string_param = 1;</code>
   */
  public com.google.protobuf.ByteString
      getStringParamBytes() {
    java.lang.Object ref = stringParam_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      stringParam_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!getStringParamBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, stringParam_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getStringParamBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, stringParam_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput)) {
      return super.equals(obj);
    }
    com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput other = (com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput) obj;

    if (!getStringParam()
        .equals(other.getStringParam())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + STRING_PARAM_FIELD_NUMBER;
    hash = (53 * hash) + getStringParam().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code codegen.DoingSomethingWithoutOutputAndActorUsecaseInput}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:codegen.DoingSomethingWithoutOutputAndActorUsecaseInput)
      com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInputOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.github.hori-ryota/go-codegen/api/example/codegen.CodegenProto.internal_static_codegen_DoingSomethingWithoutOutputAndActorUsecaseInput_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.github.hori-ryota/go-codegen/api/example/codegen.CodegenProto.internal_static_codegen_DoingSomethingWithoutOutputAndActorUsecaseInput_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput.class, com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput.Builder.class);
    }

    // Construct using com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      stringParam_ = "";

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.github.hori-ryota/go-codegen/api/example/codegen.CodegenProto.internal_static_codegen_DoingSomethingWithoutOutputAndActorUsecaseInput_descriptor;
    }

    @java.lang.Override
    public com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput getDefaultInstanceForType() {
      return com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput.getDefaultInstance();
    }

    @java.lang.Override
    public com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput build() {
      com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput buildPartial() {
      com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput result = new com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput(this);
      result.stringParam_ = stringParam_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput) {
        return mergeFrom((com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput other) {
      if (other == com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput.getDefaultInstance()) return this;
      if (!other.getStringParam().isEmpty()) {
        stringParam_ = other.stringParam_;
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private java.lang.Object stringParam_ = "";
    /**
     * <code>string string_param = 1;</code>
     */
    public java.lang.String getStringParam() {
      java.lang.Object ref = stringParam_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        stringParam_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string string_param = 1;</code>
     */
    public com.google.protobuf.ByteString
        getStringParamBytes() {
      java.lang.Object ref = stringParam_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        stringParam_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string string_param = 1;</code>
     */
    public Builder setStringParam(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      stringParam_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string string_param = 1;</code>
     */
    public Builder clearStringParam() {
      
      stringParam_ = getDefaultInstance().getStringParam();
      onChanged();
      return this;
    }
    /**
     * <code>string string_param = 1;</code>
     */
    public Builder setStringParamBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      stringParam_ = value;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:codegen.DoingSomethingWithoutOutputAndActorUsecaseInput)
  }

  // @@protoc_insertion_point(class_scope:codegen.DoingSomethingWithoutOutputAndActorUsecaseInput)
  private static final com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput();
  }

  public static com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<DoingSomethingWithoutOutputAndActorUsecaseInput>
      PARSER = new com.google.protobuf.AbstractParser<DoingSomethingWithoutOutputAndActorUsecaseInput>() {
    @java.lang.Override
    public DoingSomethingWithoutOutputAndActorUsecaseInput parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new DoingSomethingWithoutOutputAndActorUsecaseInput(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<DoingSomethingWithoutOutputAndActorUsecaseInput> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<DoingSomethingWithoutOutputAndActorUsecaseInput> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.github.hori-ryota/go-codegen/api/example/codegen.DoingSomethingWithoutOutputAndActorUsecaseInput getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

