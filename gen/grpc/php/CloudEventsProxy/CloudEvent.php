<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace CloudEventsProxy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>cloudEventsProxy.CloudEvent</code>
 */
class CloudEvent extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1 [(.validator.field) = {</code>
     */
    private $id = '';
    /**
     * Generated from protobuf field <code>string source = 2 [(.validator.field) = {</code>
     */
    private $source = '';
    /**
     * Generated from protobuf field <code>string type = 3 [(.validator.field) = {</code>
     */
    private $type = '';
    /**
     * Generated from protobuf field <code>string subject = 4;</code>
     */
    private $subject = '';
    /**
     * Generated from protobuf field <code>.google.protobuf.Struct attributes = 5;</code>
     */
    private $attributes = null;
    /**
     * Generated from protobuf field <code>.google.protobuf.Struct data = 6 [(.validator.field) = {</code>
     */
    private $data = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $source
     *     @type string $type
     *     @type string $subject
     *     @type \Google\Protobuf\Struct $attributes
     *     @type \Google\Protobuf\Struct $data
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string id = 1 [(.validator.field) = {</code>
     * @return string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>string id = 1 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string source = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getSource()
    {
        return $this->source;
    }

    /**
     * Generated from protobuf field <code>string source = 2 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setSource($var)
    {
        GPBUtil::checkString($var, True);
        $this->source = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string type = 3 [(.validator.field) = {</code>
     * @return string
     */
    public function getType()
    {
        return $this->type;
    }

    /**
     * Generated from protobuf field <code>string type = 3 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setType($var)
    {
        GPBUtil::checkString($var, True);
        $this->type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string subject = 4;</code>
     * @return string
     */
    public function getSubject()
    {
        return $this->subject;
    }

    /**
     * Generated from protobuf field <code>string subject = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setSubject($var)
    {
        GPBUtil::checkString($var, True);
        $this->subject = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Struct attributes = 5;</code>
     * @return \Google\Protobuf\Struct
     */
    public function getAttributes()
    {
        return $this->attributes;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Struct attributes = 5;</code>
     * @param \Google\Protobuf\Struct $var
     * @return $this
     */
    public function setAttributes($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Struct::class);
        $this->attributes = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Struct data = 6 [(.validator.field) = {</code>
     * @return \Google\Protobuf\Struct
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Struct data = 6 [(.validator.field) = {</code>
     * @param \Google\Protobuf\Struct $var
     * @return $this
     */
    public function setData($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Struct::class);
        $this->data = $var;

        return $this;
    }

}

