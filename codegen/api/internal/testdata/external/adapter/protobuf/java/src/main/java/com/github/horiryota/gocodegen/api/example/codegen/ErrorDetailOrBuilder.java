// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: error.proto

package com.github.horiryota.gocodegen.api.example.codegen;

public interface ErrorDetailOrBuilder extends
    // @@protoc_insertion_point(interface_extends:codegen.ErrorDetail)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string code = 1;</code>
   */
  java.lang.String getCode();
  /**
   * <code>string code = 1;</code>
   */
  com.google.protobuf.ByteString
      getCodeBytes();

  /**
   * <code>repeated string args = 2;</code>
   */
  java.util.List<java.lang.String>
      getArgsList();
  /**
   * <code>repeated string args = 2;</code>
   */
  int getArgsCount();
  /**
   * <code>repeated string args = 2;</code>
   */
  java.lang.String getArgs(int index);
  /**
   * <code>repeated string args = 2;</code>
   */
  com.google.protobuf.ByteString
      getArgsBytes(int index);
}