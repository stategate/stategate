<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Eventgate;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Object hold's the current state of an object
 *
 * Generated from protobuf message <code>eventgate.Object</code>
 */
class Object extends \Google\Protobuf\Internal\Message
{
    /**
     * Object type (ex: user)
     *
     * Generated from protobuf field <code>string type = 1 [(.validator.field) = {</code>
     */
    private $type = '';
    /**
     * Object key (unique within type)
     *
     * Generated from protobuf field <code>string key = 2 [(.validator.field) = {</code>
     */
    private $key = '';
    /**
     * Object values (structured k/v pairs)
     *
     * Generated from protobuf field <code>.google.protobuf.Struct values = 3 [(.validator.field) = {</code>
     */
    private $values = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $type
     *           Object type (ex: user)
     *     @type string $key
     *           Object key (unique within type)
     *     @type \Google\Protobuf\Struct $values
     *           Object values (structured k/v pairs)
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Object type (ex: user)
     *
     * Generated from protobuf field <code>string type = 1 [(.validator.field) = {</code>
     * @return string
     */
    public function getType()
    {
        return $this->type;
    }

    /**
     * Object type (ex: user)
     *
     * Generated from protobuf field <code>string type = 1 [(.validator.field) = {</code>
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
     * Object key (unique within type)
     *
     * Generated from protobuf field <code>string key = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getKey()
    {
        return $this->key;
    }

    /**
     * Object key (unique within type)
     *
     * Generated from protobuf field <code>string key = 2 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->key = $var;

        return $this;
    }

    /**
     * Object values (structured k/v pairs)
     *
     * Generated from protobuf field <code>.google.protobuf.Struct values = 3 [(.validator.field) = {</code>
     * @return \Google\Protobuf\Struct
     */
    public function getValues()
    {
        return $this->values;
    }

    /**
     * Object values (structured k/v pairs)
     *
     * Generated from protobuf field <code>.google.protobuf.Struct values = 3 [(.validator.field) = {</code>
     * @param \Google\Protobuf\Struct $var
     * @return $this
     */
    public function setValues($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Struct::class);
        $this->values = $var;

        return $this;
    }

}
