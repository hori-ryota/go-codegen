// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: error.proto

package com.github.horiryota.gocodegen.api.example.codegen;

/**
 * Protobuf type {@code codegen.Error}
 */
public  final class Error extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:codegen.Error)
    ErrorOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Error.newBuilder() to construct.
  private Error(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Error() {
    code_ = "";
    details_ = java.util.Collections.emptyList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Error();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Error(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
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

            code_ = s;
            break;
          }
          case 18: {
            if (!((mutable_bitField0_ & 0x00000001) != 0)) {
              details_ = new java.util.ArrayList<com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail>();
              mutable_bitField0_ |= 0x00000001;
            }
            details_.add(
                input.readMessage(com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.parser(), extensionRegistry));
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
      if (((mutable_bitField0_ & 0x00000001) != 0)) {
        details_ = java.util.Collections.unmodifiableList(details_);
      }
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.github.horiryota.gocodegen.api.example.codegen.ErrorProto.internal_static_codegen_Error_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.github.horiryota.gocodegen.api.example.codegen.ErrorProto.internal_static_codegen_Error_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.github.horiryota.gocodegen.api.example.codegen.Error.class, com.github.horiryota.gocodegen.api.example.codegen.Error.Builder.class);
  }

  public static final int CODE_FIELD_NUMBER = 1;
  private volatile java.lang.Object code_;
  /**
   * <code>string code = 1;</code>
   */
  public java.lang.String getCode() {
    java.lang.Object ref = code_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      code_ = s;
      return s;
    }
  }
  /**
   * <code>string code = 1;</code>
   */
  public com.google.protobuf.ByteString
      getCodeBytes() {
    java.lang.Object ref = code_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      code_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int DETAILS_FIELD_NUMBER = 2;
  private java.util.List<com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail> details_;
  /**
   * <code>repeated .codegen.ErrorDetail details = 2;</code>
   */
  public java.util.List<com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail> getDetailsList() {
    return details_;
  }
  /**
   * <code>repeated .codegen.ErrorDetail details = 2;</code>
   */
  public java.util.List<? extends com.github.horiryota.gocodegen.api.example.codegen.ErrorDetailOrBuilder> 
      getDetailsOrBuilderList() {
    return details_;
  }
  /**
   * <code>repeated .codegen.ErrorDetail details = 2;</code>
   */
  public int getDetailsCount() {
    return details_.size();
  }
  /**
   * <code>repeated .codegen.ErrorDetail details = 2;</code>
   */
  public com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail getDetails(int index) {
    return details_.get(index);
  }
  /**
   * <code>repeated .codegen.ErrorDetail details = 2;</code>
   */
  public com.github.horiryota.gocodegen.api.example.codegen.ErrorDetailOrBuilder getDetailsOrBuilder(
      int index) {
    return details_.get(index);
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
    if (!getCodeBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, code_);
    }
    for (int i = 0; i < details_.size(); i++) {
      output.writeMessage(2, details_.get(i));
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getCodeBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, code_);
    }
    for (int i = 0; i < details_.size(); i++) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, details_.get(i));
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
    if (!(obj instanceof com.github.horiryota.gocodegen.api.example.codegen.Error)) {
      return super.equals(obj);
    }
    com.github.horiryota.gocodegen.api.example.codegen.Error other = (com.github.horiryota.gocodegen.api.example.codegen.Error) obj;

    if (!getCode()
        .equals(other.getCode())) return false;
    if (!getDetailsList()
        .equals(other.getDetailsList())) return false;
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
    hash = (37 * hash) + CODE_FIELD_NUMBER;
    hash = (53 * hash) + getCode().hashCode();
    if (getDetailsCount() > 0) {
      hash = (37 * hash) + DETAILS_FIELD_NUMBER;
      hash = (53 * hash) + getDetailsList().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.horiryota.gocodegen.api.example.codegen.Error parseFrom(
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
  public static Builder newBuilder(com.github.horiryota.gocodegen.api.example.codegen.Error prototype) {
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
   * Protobuf type {@code codegen.Error}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:codegen.Error)
      com.github.horiryota.gocodegen.api.example.codegen.ErrorOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.github.horiryota.gocodegen.api.example.codegen.ErrorProto.internal_static_codegen_Error_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.github.horiryota.gocodegen.api.example.codegen.ErrorProto.internal_static_codegen_Error_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.github.horiryota.gocodegen.api.example.codegen.Error.class, com.github.horiryota.gocodegen.api.example.codegen.Error.Builder.class);
    }

    // Construct using com.github.horiryota.gocodegen.api.example.codegen.Error.newBuilder()
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
        getDetailsFieldBuilder();
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      code_ = "";

      if (detailsBuilder_ == null) {
        details_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
      } else {
        detailsBuilder_.clear();
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.github.horiryota.gocodegen.api.example.codegen.ErrorProto.internal_static_codegen_Error_descriptor;
    }

    @java.lang.Override
    public com.github.horiryota.gocodegen.api.example.codegen.Error getDefaultInstanceForType() {
      return com.github.horiryota.gocodegen.api.example.codegen.Error.getDefaultInstance();
    }

    @java.lang.Override
    public com.github.horiryota.gocodegen.api.example.codegen.Error build() {
      com.github.horiryota.gocodegen.api.example.codegen.Error result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.github.horiryota.gocodegen.api.example.codegen.Error buildPartial() {
      com.github.horiryota.gocodegen.api.example.codegen.Error result = new com.github.horiryota.gocodegen.api.example.codegen.Error(this);
      int from_bitField0_ = bitField0_;
      result.code_ = code_;
      if (detailsBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0)) {
          details_ = java.util.Collections.unmodifiableList(details_);
          bitField0_ = (bitField0_ & ~0x00000001);
        }
        result.details_ = details_;
      } else {
        result.details_ = detailsBuilder_.build();
      }
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
      if (other instanceof com.github.horiryota.gocodegen.api.example.codegen.Error) {
        return mergeFrom((com.github.horiryota.gocodegen.api.example.codegen.Error)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.github.horiryota.gocodegen.api.example.codegen.Error other) {
      if (other == com.github.horiryota.gocodegen.api.example.codegen.Error.getDefaultInstance()) return this;
      if (!other.getCode().isEmpty()) {
        code_ = other.code_;
        onChanged();
      }
      if (detailsBuilder_ == null) {
        if (!other.details_.isEmpty()) {
          if (details_.isEmpty()) {
            details_ = other.details_;
            bitField0_ = (bitField0_ & ~0x00000001);
          } else {
            ensureDetailsIsMutable();
            details_.addAll(other.details_);
          }
          onChanged();
        }
      } else {
        if (!other.details_.isEmpty()) {
          if (detailsBuilder_.isEmpty()) {
            detailsBuilder_.dispose();
            detailsBuilder_ = null;
            details_ = other.details_;
            bitField0_ = (bitField0_ & ~0x00000001);
            detailsBuilder_ = 
              com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                 getDetailsFieldBuilder() : null;
          } else {
            detailsBuilder_.addAllMessages(other.details_);
          }
        }
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
      com.github.horiryota.gocodegen.api.example.codegen.Error parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (com.github.horiryota.gocodegen.api.example.codegen.Error) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private java.lang.Object code_ = "";
    /**
     * <code>string code = 1;</code>
     */
    public java.lang.String getCode() {
      java.lang.Object ref = code_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        code_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string code = 1;</code>
     */
    public com.google.protobuf.ByteString
        getCodeBytes() {
      java.lang.Object ref = code_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        code_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string code = 1;</code>
     */
    public Builder setCode(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      code_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string code = 1;</code>
     */
    public Builder clearCode() {
      
      code_ = getDefaultInstance().getCode();
      onChanged();
      return this;
    }
    /**
     * <code>string code = 1;</code>
     */
    public Builder setCodeBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      code_ = value;
      onChanged();
      return this;
    }

    private java.util.List<com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail> details_ =
      java.util.Collections.emptyList();
    private void ensureDetailsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        details_ = new java.util.ArrayList<com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail>(details_);
        bitField0_ |= 0x00000001;
       }
    }

    private com.google.protobuf.RepeatedFieldBuilderV3<
        com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetailOrBuilder> detailsBuilder_;

    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public java.util.List<com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail> getDetailsList() {
      if (detailsBuilder_ == null) {
        return java.util.Collections.unmodifiableList(details_);
      } else {
        return detailsBuilder_.getMessageList();
      }
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public int getDetailsCount() {
      if (detailsBuilder_ == null) {
        return details_.size();
      } else {
        return detailsBuilder_.getCount();
      }
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail getDetails(int index) {
      if (detailsBuilder_ == null) {
        return details_.get(index);
      } else {
        return detailsBuilder_.getMessage(index);
      }
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public Builder setDetails(
        int index, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail value) {
      if (detailsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureDetailsIsMutable();
        details_.set(index, value);
        onChanged();
      } else {
        detailsBuilder_.setMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public Builder setDetails(
        int index, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder builderForValue) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        details_.set(index, builderForValue.build());
        onChanged();
      } else {
        detailsBuilder_.setMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public Builder addDetails(com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail value) {
      if (detailsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureDetailsIsMutable();
        details_.add(value);
        onChanged();
      } else {
        detailsBuilder_.addMessage(value);
      }
      return this;
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public Builder addDetails(
        int index, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail value) {
      if (detailsBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        ensureDetailsIsMutable();
        details_.add(index, value);
        onChanged();
      } else {
        detailsBuilder_.addMessage(index, value);
      }
      return this;
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public Builder addDetails(
        com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder builderForValue) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        details_.add(builderForValue.build());
        onChanged();
      } else {
        detailsBuilder_.addMessage(builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public Builder addDetails(
        int index, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder builderForValue) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        details_.add(index, builderForValue.build());
        onChanged();
      } else {
        detailsBuilder_.addMessage(index, builderForValue.build());
      }
      return this;
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public Builder addAllDetails(
        java.lang.Iterable<? extends com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail> values) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        com.google.protobuf.AbstractMessageLite.Builder.addAll(
            values, details_);
        onChanged();
      } else {
        detailsBuilder_.addAllMessages(values);
      }
      return this;
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public Builder clearDetails() {
      if (detailsBuilder_ == null) {
        details_ = java.util.Collections.emptyList();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
      } else {
        detailsBuilder_.clear();
      }
      return this;
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public Builder removeDetails(int index) {
      if (detailsBuilder_ == null) {
        ensureDetailsIsMutable();
        details_.remove(index);
        onChanged();
      } else {
        detailsBuilder_.remove(index);
      }
      return this;
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder getDetailsBuilder(
        int index) {
      return getDetailsFieldBuilder().getBuilder(index);
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public com.github.horiryota.gocodegen.api.example.codegen.ErrorDetailOrBuilder getDetailsOrBuilder(
        int index) {
      if (detailsBuilder_ == null) {
        return details_.get(index);  } else {
        return detailsBuilder_.getMessageOrBuilder(index);
      }
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public java.util.List<? extends com.github.horiryota.gocodegen.api.example.codegen.ErrorDetailOrBuilder> 
         getDetailsOrBuilderList() {
      if (detailsBuilder_ != null) {
        return detailsBuilder_.getMessageOrBuilderList();
      } else {
        return java.util.Collections.unmodifiableList(details_);
      }
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder addDetailsBuilder() {
      return getDetailsFieldBuilder().addBuilder(
          com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.getDefaultInstance());
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder addDetailsBuilder(
        int index) {
      return getDetailsFieldBuilder().addBuilder(
          index, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.getDefaultInstance());
    }
    /**
     * <code>repeated .codegen.ErrorDetail details = 2;</code>
     */
    public java.util.List<com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder> 
         getDetailsBuilderList() {
      return getDetailsFieldBuilder().getBuilderList();
    }
    private com.google.protobuf.RepeatedFieldBuilderV3<
        com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetailOrBuilder> 
        getDetailsFieldBuilder() {
      if (detailsBuilder_ == null) {
        detailsBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
            com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetail.Builder, com.github.horiryota.gocodegen.api.example.codegen.ErrorDetailOrBuilder>(
                details_,
                ((bitField0_ & 0x00000001) != 0),
                getParentForChildren(),
                isClean());
        details_ = null;
      }
      return detailsBuilder_;
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


    // @@protoc_insertion_point(builder_scope:codegen.Error)
  }

  // @@protoc_insertion_point(class_scope:codegen.Error)
  private static final com.github.horiryota.gocodegen.api.example.codegen.Error DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.github.horiryota.gocodegen.api.example.codegen.Error();
  }

  public static com.github.horiryota.gocodegen.api.example.codegen.Error getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Error>
      PARSER = new com.google.protobuf.AbstractParser<Error>() {
    @java.lang.Override
    public Error parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Error(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Error> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Error> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.github.horiryota.gocodegen.api.example.codegen.Error getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

