// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: schema.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Eventgate {

  /// <summary>Holder for reflection information generated from schema.proto</summary>
  public static partial class SchemaReflection {

    #region Descriptor
    /// <summary>File descriptor for schema.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static SchemaReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CgxzY2hlbWEucHJvdG8SCWV2ZW50Z2F0ZRocZ29vZ2xlL2FwaS9hbm5vdGF0",
            "aW9ucy5wcm90bxocZ29vZ2xlL3Byb3RvYnVmL3N0cnVjdC5wcm90bxofZ29v",
            "Z2xlL3Byb3RvYnVmL3RpbWVzdGFtcC5wcm90bxoZZ29vZ2xlL3Byb3RvYnVm",
            "L2FueS5wcm90bxobZ29vZ2xlL3Byb3RvYnVmL2VtcHR5LnByb3RvGjZnaXRo",
            "dWIuY29tL213aXRrb3cvZ28tcHJvdG8tdmFsaWRhdG9ycy92YWxpZGF0b3Iu",
            "cHJvdG8iPgoLUmVjZWl2ZU9wdHMSFwoHY2hhbm5lbBgBIAEoCUIG4t8fAlgB",
            "EhYKDmNvbnN1bWVyX2dyb3VwGAIgASgJIugBCgVFdmVudBIKCgJpZBgBIAEo",
            "CRIXCgdjaGFubmVsGAIgASgJQgbi3x8CWAESLQoEZGF0YRgFIAEoCzIXLmdv",
            "b2dsZS5wcm90b2J1Zi5TdHJ1Y3RCBuLfHwIgARIwCghtZXRhZGF0YRgGIAMo",
            "CzIeLmV2ZW50Z2F0ZS5FdmVudC5NZXRhZGF0YUVudHJ5EigKBHRpbWUYFCAB",
            "KAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wGi8KDU1ldGFkYXRhRW50",
            "cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgJOgI4ATKfAQoQRXZlbnRH",
            "YXRlU2VydmljZRJCCgRTZW5kEhAuZXZlbnRnYXRlLkV2ZW50GhYuZ29vZ2xl",
            "LnByb3RvYnVmLkVtcHR5IhCC0+STAgoaBS9zZW5kOgEqEkcKB1JlY2VpdmUS",
            "Fi5ldmVudGdhdGUuUmVjZWl2ZU9wdHMaEC5ldmVudGdhdGUuRXZlbnQiEILT",
            "5JMCChIIL3JlY2VpdmUwAUILWglldmVudGdhdGViBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Api.AnnotationsReflection.Descriptor, global::Google.Protobuf.WellKnownTypes.StructReflection.Descriptor, global::Google.Protobuf.WellKnownTypes.TimestampReflection.Descriptor, global::Google.Protobuf.WellKnownTypes.AnyReflection.Descriptor, global::Google.Protobuf.WellKnownTypes.EmptyReflection.Descriptor, global::Validator.ValidatorReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Eventgate.ReceiveOpts), global::Eventgate.ReceiveOpts.Parser, new[]{ "Channel", "ConsumerGroup" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Eventgate.Event), global::Eventgate.Event.Parser, new[]{ "Id", "Channel", "Data", "Metadata", "Time" }, null, null, new pbr::GeneratedClrTypeInfo[] { null, })
          }));
    }
    #endregion

  }
  #region Messages
  /// <summary>
  /// ReceiveOpts filters events before they are received by a consumer
  /// </summary>
  public sealed partial class ReceiveOpts : pb::IMessage<ReceiveOpts> {
    private static readonly pb::MessageParser<ReceiveOpts> _parser = new pb::MessageParser<ReceiveOpts>(() => new ReceiveOpts());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<ReceiveOpts> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Eventgate.SchemaReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ReceiveOpts() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ReceiveOpts(ReceiveOpts other) : this() {
      channel_ = other.channel_;
      consumerGroup_ = other.consumerGroup_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public ReceiveOpts Clone() {
      return new ReceiveOpts(this);
    }

    /// <summary>Field number for the "channel" field.</summary>
    public const int ChannelFieldNumber = 1;
    private string channel_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Channel {
      get { return channel_; }
      set {
        channel_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "consumer_group" field.</summary>
    public const int ConsumerGroupFieldNumber = 2;
    private string consumerGroup_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string ConsumerGroup {
      get { return consumerGroup_; }
      set {
        consumerGroup_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as ReceiveOpts);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(ReceiveOpts other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Channel != other.Channel) return false;
      if (ConsumerGroup != other.ConsumerGroup) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Channel.Length != 0) hash ^= Channel.GetHashCode();
      if (ConsumerGroup.Length != 0) hash ^= ConsumerGroup.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Channel.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Channel);
      }
      if (ConsumerGroup.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(ConsumerGroup);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Channel.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Channel);
      }
      if (ConsumerGroup.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(ConsumerGroup);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(ReceiveOpts other) {
      if (other == null) {
        return;
      }
      if (other.Channel.Length != 0) {
        Channel = other.Channel;
      }
      if (other.ConsumerGroup.Length != 0) {
        ConsumerGroup = other.ConsumerGroup;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Channel = input.ReadString();
            break;
          }
          case 18: {
            ConsumerGroup = input.ReadString();
            break;
          }
        }
      }
    }

  }

  /// <summary>
  /// Event is a specification for describing event data
  /// </summary>
  public sealed partial class Event : pb::IMessage<Event> {
    private static readonly pb::MessageParser<Event> _parser = new pb::MessageParser<Event>(() => new Event());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Event> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Eventgate.SchemaReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Event() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Event(Event other) : this() {
      id_ = other.id_;
      channel_ = other.channel_;
      data_ = other.data_ != null ? other.data_.Clone() : null;
      metadata_ = other.metadata_.Clone();
      time_ = other.time_ != null ? other.time_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Event Clone() {
      return new Event(this);
    }

    /// <summary>Field number for the "id" field.</summary>
    public const int IdFieldNumber = 1;
    private string id_ = "";
    /// <summary>
    /// Identifies the event. If an ID is not sent with the event, a uuid will be assigned
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Id {
      get { return id_; }
      set {
        id_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "channel" field.</summary>
    public const int ChannelFieldNumber = 2;
    private string channel_ = "";
    /// <summary>
    /// Identifies the channel/subject to which the event will be sent
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Channel {
      get { return channel_; }
      set {
        channel_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "data" field.</summary>
    public const int DataFieldNumber = 5;
    private global::Google.Protobuf.WellKnownTypes.Struct data_;
    /// <summary>
    /// The event payload(structured).
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Google.Protobuf.WellKnownTypes.Struct Data {
      get { return data_; }
      set {
        data_ = value;
      }
    }

    /// <summary>Field number for the "metadata" field.</summary>
    public const int MetadataFieldNumber = 6;
    private static readonly pbc::MapField<string, string>.Codec _map_metadata_codec
        = new pbc::MapField<string, string>.Codec(pb::FieldCodec.ForString(10), pb::FieldCodec.ForString(18), 50);
    private readonly pbc::MapField<string, string> metadata_ = new pbc::MapField<string, string>();
    /// <summary>
    /// Arbitrary metadata about the event
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::MapField<string, string> Metadata {
      get { return metadata_; }
    }

    /// <summary>Field number for the "time" field.</summary>
    public const int TimeFieldNumber = 20;
    private global::Google.Protobuf.WellKnownTypes.Timestamp time_;
    /// <summary>
    /// Timestamp of when the occurrence happened. Must adhere to RFC 3339. If a timestamp is not sent with the event, the current time will be assigned
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Google.Protobuf.WellKnownTypes.Timestamp Time {
      get { return time_; }
      set {
        time_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Event);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Event other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Id != other.Id) return false;
      if (Channel != other.Channel) return false;
      if (!object.Equals(Data, other.Data)) return false;
      if (!Metadata.Equals(other.Metadata)) return false;
      if (!object.Equals(Time, other.Time)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Id.Length != 0) hash ^= Id.GetHashCode();
      if (Channel.Length != 0) hash ^= Channel.GetHashCode();
      if (data_ != null) hash ^= Data.GetHashCode();
      hash ^= Metadata.GetHashCode();
      if (time_ != null) hash ^= Time.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Id.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Id);
      }
      if (Channel.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(Channel);
      }
      if (data_ != null) {
        output.WriteRawTag(42);
        output.WriteMessage(Data);
      }
      metadata_.WriteTo(output, _map_metadata_codec);
      if (time_ != null) {
        output.WriteRawTag(162, 1);
        output.WriteMessage(Time);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Id.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Id);
      }
      if (Channel.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Channel);
      }
      if (data_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Data);
      }
      size += metadata_.CalculateSize(_map_metadata_codec);
      if (time_ != null) {
        size += 2 + pb::CodedOutputStream.ComputeMessageSize(Time);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Event other) {
      if (other == null) {
        return;
      }
      if (other.Id.Length != 0) {
        Id = other.Id;
      }
      if (other.Channel.Length != 0) {
        Channel = other.Channel;
      }
      if (other.data_ != null) {
        if (data_ == null) {
          data_ = new global::Google.Protobuf.WellKnownTypes.Struct();
        }
        Data.MergeFrom(other.Data);
      }
      metadata_.Add(other.metadata_);
      if (other.time_ != null) {
        if (time_ == null) {
          time_ = new global::Google.Protobuf.WellKnownTypes.Timestamp();
        }
        Time.MergeFrom(other.Time);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Id = input.ReadString();
            break;
          }
          case 18: {
            Channel = input.ReadString();
            break;
          }
          case 42: {
            if (data_ == null) {
              data_ = new global::Google.Protobuf.WellKnownTypes.Struct();
            }
            input.ReadMessage(data_);
            break;
          }
          case 50: {
            metadata_.AddEntriesFrom(input, _map_metadata_codec);
            break;
          }
          case 162: {
            if (time_ == null) {
              time_ = new global::Google.Protobuf.WellKnownTypes.Timestamp();
            }
            input.ReadMessage(time_);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
